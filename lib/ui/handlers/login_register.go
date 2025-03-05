package handlers


import(
  "fmt"
  "time"
  //"context"
  "net/http"
  dom"ken/lib/domain"
  ent"ken/lib/entities"
  "ken/lib/utils"
)

func (hnd *Handler) Login(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
    req.ParseForm()
    pass := req.FormValue("password")
    email := req.FormValue("email")
    fmt.Println("Pass: ",pass)
    fmt.Println("Email: ",email)
    /*conn,err := hnd.Dbs.Conn(context.Context())
    if err != nil {
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error getting connection: %s",err))
      utils.Logerror(e)
      // redirect or something
      hnd.Internalserverror(res,req)
      return
    }*/
    dmn := &dom.Domain{Dbs:hnd.Dbs}
    ud,isAuth := dmn.Authenticate(email,pass)
    if !isAuth {
      fmt.Println("Failed login attempt")
      tpl,err := hnd.GetATemplate("login","login.tmpl")
      if err != nil {
        utils.Warning(fmt.Sprintf("%s",err))
        hnd.Internalserverror(res,req)
        //http.Error(res, "An error occurred", http.StatusInternalServerError)
        return
      }
      fmt.Println("Error wrong password..")
      tpl.ExecuteTemplate(res,"login","Wrong password or email. Try again!")
      return
    }
    //set session
    token,err := hnd.GenerateJWT(ud)
    if err != nil {
      utils.Warning(fmt.Sprintf("Error generating JWT: %s",err))
      hnd.Internalserverror(res,req)
      return
    }
    res.Header().Set("Authorization", token)
    cookie := http.Cookie{
			Name:     "Authorization",
			Value:    token,
			//HttpOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour), // Set an appropriate expiration time
		}
		http.SetCookie(res, &cookie)
    //redirect to dashboard or get the dash data and execute dash
    http.Redirect(res,req,"/home",http.StatusSeeOther)
    return
  }
  tpl,err := hnd.GetATemplate("login","login.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
    return
  }
  tpl.ExecuteTemplate(res,"login",nil)
  return
}

func (hnd *Handler) Logout(res http.ResponseWriter, req *http.Request){
  cookie,err := req.Cookie("Authorization")
  if err == http.ErrNoCookie {
    http.Redirect(res,req,"/login",http.StatusSeeOther)
    return
  } else if err != nil {
      utils.Warning(fmt.Sprintf("[+]  Some internal error. \nERROR: ",err))
      hnd.Internalserverror(res,req)
      return
  }
  tokenString := cookie.Value
  req.Header.Del("Authorization")
  res.Header().Del("Authorization")
  hnd.InvalidTokens = append(hnd.InvalidTokens,tokenString)
  tpl,err := hnd.GetATemplate("login","login.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
    return
  }
  tpl.ExecuteTemplate(res,"login","Logged Out. ADIOS!!!")
  return
}


func (hnd *Handler) Register(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST" {
    //do the post thing
    req.ParseForm()
    pass := req.FormValue("password")
    email := req.FormValue("email")
    name := req.FormValue("name")
    phone := req.FormValue("phone")
    dmn := &dom.Domain{Dbs:hnd.Dbs}
    ud := ent.UserData {
      UserID: utils.GenerateUUID(),
      Role: "ADMIN",
      Phone: phone,
      Name: name,
      Email: email,
      Password: pass,
    }
    ud.Touch()
    err := dmn.CreateUser(ud)
    if err != nil {
      utils.Warning(fmt.Sprintf("Error registering user: %s",err))
      hnd.Internalserverror(res,req)
      return
    }
    http.Redirect(res,req,"/login",http.StatusSeeOther)
    return
  }
  tpl,err := hnd.GetATemplate("register","sighnup.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
    return
  }
  tpl.ExecuteTemplate(res,"register",nil)
  return
}
