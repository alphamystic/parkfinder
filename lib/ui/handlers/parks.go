package handlers


import(
  "io"
  "os"
  "fmt"
  "io/fs"
  //"context"
  "net/http"
  "path/filepath"
  "encoding/json"
  dom"ken/lib/domain"
  ent"ken/lib/entities"
  "ken/lib/utils"
)

// list parks
func (hnd *Handler) Parks(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("parks","parks.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  dmn := &dom.Domain{Dbs:hnd.Dbs}
  parks,err := dmn.ListParks()
  if err != nil {
    utils.Warning(fmt.Sprintf("Error getting parks: %s",err))
    return
  }
  tpl.ExecuteTemplate(res,"parks",parks)
  return
}

func (hnd *Handler) Viewpark(res http.ResponseWriter, req *http.Request) {
  query := req.URL.Query()
  pid := query.Get("pid")
  dmn := &dom.Domain{Dbs:hnd.Dbs}
  park,err := dmn.ViewPark(pid)
  if err != nil {
    utils.Warning(fmt.Sprintf("Error viewing park: %s",err))
    hnd.Internalserverror(res,req)
    return
  }
  reviews,err := dmn.ListReviews(park.ParkID)
  if err != nil {
    utils.Warning(fmt.Sprintf("Error getting reviews: %s",err))
  }
  tpl,err := hnd.GetATemplate("view_park","view_park.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    //http.Error(res, "An error occurred", http.StatusInternalServerError)
    hnd.Internalserverror(res,req)
    return
  }
  images,err := getImageFiles("./lib/ui/static/uploads/"+park.ParkID,park.ParkID)
  if err != nil{
    utils.Warning(fmt.Sprintf("Error getting images for park with id: %s",park.ParkID))
  }
  fmt.Println("lengt images: ",len(images))
  tpl.ExecuteTemplate(res,"view_park",HOME{
    "Park": park,
    "Reviews": reviews,
    "Images":images,
  })
  return
}

type Image struct{
  Dir string
  Image string
}

func getImageFiles(dir,pid string) ([]Image, error) {
	var images []Image
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is an image
		if !d.IsDir() {
			ext := filepath.Ext(d.Name())
			if isImageFile(ext) {
				image := Image{
					Dir:   pid,//filepath.Dir(path),
					Image: d.Name(),
				}
				images = append(images, image)
			}
		}
		return nil
	})

	return images, err
}

// isImageFile checks if a file extension corresponds to an image file.
func isImageFile(ext string) bool {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp":
		return true
	}
	return false
}



func (hnd *Handler) CreateParks(res http.ResponseWriter, req *http.Request) {
	ud, err := hnd.GetUDFromToken(req)
	if err != nil {
		utils.Warning(fmt.Sprintf("%s", err))
		http.Redirect(res, req, "/logout", http.StatusSeeOther)
		return
	}
  fmt.Println("Role: ",ud.Role)
	if ud.Role != "ADMIN" {
    utils.Warning("None admin tryig to create park...")
		http.Redirect(res, req, "/logout", http.StatusSeeOther)
		return
	}
	if req.Method == "POST" {
		err := req.ParseMultipartForm(50 << 20) // 50MB limit
		if err != nil {
			utils.Warning(fmt.Sprintf("Error parsing form: %s", err))
			hnd.Internalserverror(res, req)
			return
		}
		req.ParseForm()
		name := req.FormValue("park_name")
		loc := req.FormValue("location")
		description := req.FormValue("description")
		d := ent.Park{
			ParkID:      utils.GenerateUUID(),
			Name:        name,
			Location:    loc,
			Description: description,
		}
		d.Touch()
		dmn := &dom.Domain{Dbs: hnd.Dbs}
		if err := dmn.CreatePark(d); err != nil {
			utils.Warning(fmt.Sprintf("%s", err))
			hnd.Internalserverror(res, req)
			return
		}
		var imagePaths []string
		files := req.MultipartForm.File["images"]

		// Create the directory to save the uploaded images
		dir := fmt.Sprintf("./lib/ui/static/uploads/%s", d.ParkID)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			utils.Warning(fmt.Sprintf("Error creating directory: %s", err))
			hnd.Internalserverror(res, req)
			return
		}
		for _, fileHeader := range files {
			// Open the uploaded file
			file, err := fileHeader.Open()
			if err != nil {
				utils.Warning(fmt.Sprintf("Error opening file: %s", err))
				hnd.Internalserverror(res, req)
				return
			}
			defer file.Close()
			// Create a temporary file to save the uploaded file
			tempFile, err := os.CreateTemp(dir, fmt.Sprintf("park-%s-*%s", utils.GenerateUUID(), filepath.Ext(fileHeader.Filename)))
			if err != nil {
				utils.Warning(fmt.Sprintf("Error creating temporary file: %s", err))
				hnd.Internalserverror(res, req)
				return
			}
			defer tempFile.Close()
			// Copy the uploaded file to the temporary file
			_, err = io.Copy(tempFile, file)
			if err != nil {
				utils.Warning(fmt.Sprintf("Error saving file: %s", err))
				hnd.Internalserverror(res, req)
				return
			}
			// Add the temp file path to the list of image paths
			imagePaths = append(imagePaths, tempFile.Name())
		}
		http.Redirect(res, req, fmt.Sprintf("/viewpark?pid=%s", d.ParkID), http.StatusSeeOther)
		return
	}
	tpl, err := hnd.GetATemplate("create_parks", "create_park.tmpl")
	if err != nil {
		utils.Warning(fmt.Sprintf("%s", err))
		//http.Error(res, "An error occurred", http.StatusInternalServerError)
		hnd.Internalserverror(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "create_parks", nil)
	return
}



type PR struct {
  ParkID string `json:"parkid"`
  Comment string `json: comment`
}

func (hnd *Handler) CreatParkReview(res http.ResponseWriter, req *http.Request) {
  fmt.Println("Runing create review")
  //takes in a json and returns a success json
  var pr PR
  if err := json.NewDecoder(req.Body).Decode(&pr); err != nil {
    res.WriteHeader(http.StatusBadRequest)
		http.Error(res, "Failed to decode request body", http.StatusBadRequest)
    fmt.Println("[-] ERROR: ",err)
		return
	}
  ud,err := hnd.GetUDFromToken(req)
  if err != nil{
    http.Redirect(res,req,"/login",http.StatusSeeOther)
    return
  }
  dmn := &dom.Domain{Dbs:hnd.Dbs}
  rvw := ent.Review{
    ReviewID: utils.GenerateUUID(),
    ParkID: pr.ParkID,
    UserID: ud.UserID,
    UserName: ud.Name,
    Comment: pr.Comment,
  }
  rvw.Touch()
  if err := dmn.CreateReview(rvw); err != nil {
    utils.Warning(fmt.Sprintf("Error creating review: %s",err))
    resp := Result{
      Success: false,
      Message: "Error creating review",
    }
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusInternalServerError)
    if err := json.NewEncoder(res).Encode(resp); err != nil {
      http.Error(res, "Failed to create review", http.StatusInternalServerError)
      return
    }
    return
  }
  fmt.Println("Successfully created review")
  resp := Result{
    Success: true,
    Message: "Created review.",
  }
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(res).Encode(resp); err != nil {
    http.Error(res, "Failed to create review", http.StatusInternalServerError)
    return
  }
  return
}


type Result struct {
  Success bool
  Message string
}
