package handlers

import(
  "fmt"
  "net/http"
  "ken/lib/utils"
)

func (hnd *Handler) P404(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("nfh","404.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"nfh",nil)
}


func (hnd *Handler) Home(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("home","body.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"home",nil)
  return
}

func (hnd *Handler) Internalserverror(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("err500","500.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"err500",nil)
  return
}


func (hnd *Handler) About(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("about","about.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"about",nil)
}


func (hnd *Handler) Contact(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("contact","contact.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"contact",nil)
}
