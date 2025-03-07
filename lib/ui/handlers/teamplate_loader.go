package handlers

import (
	"fmt"
	"strings"
	"io/ioutil"
	"html/template"
)

/*
var PAGES []{string}{}

func (hnd *Handler) LoadPages(){
	// read the whole Directory
	pages,err := ioutil.ReadDir(hnd.TemplatesDir + "pages/")
	if err != nil {
		return fmt.Errorf("Error reading body Directory: %q",err)
	}
	// For each, read the pages and index it with the name.
	for _, page := range pages {
		name,data
	}
}
*/
// Creates a whole new base template for serving
func (hnd *Handler) LoadBase() error {
	baseHTML, err := ioutil.ReadFile(hnd.TemplatesDir +"base.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading base template: %q",err)
	}
	navbar, err := ioutil.ReadFile(hnd.TemplatesDir + "navbar.tmpl")
	if err != nil {
		return fmt.Errorf("Error loading sidebar template: %q",err)
	}
	footer, err := ioutil.ReadFile(hnd.TemplatesDir + "footer.tmpl")
	if err != nil {
		return  fmt.Errorf("Error loading footer template: %q",err)
	}

	// Replace placeholders in the base HTML with actual content
	combinedHTML := strings.ReplaceAll(string(baseHTML), "{{.NAVBAR}}", string(navbar))
	combinedHTML = strings.ReplaceAll(combinedHTML, "{{.FOOTER}}", string(footer))
	hnd.Base = combinedHTML
	//var tpl = new(template.Template)
	var tpl = template.New("base")
	tpl,err = tpl.Parse(string(combinedHTML))
	if err != nil {
		return fmt.Errorf("Error parsing combined html to a template: %q",err)
	}
	hnd.Tpl = tpl
	return  nil
}

func (hnd *Handler) GetATemplate(name,templFile string) (*template.Template,error) {
	body,err := ioutil.ReadFile(hnd.TemplatesDir + "pages/" + templFile)
	if err != nil {
		return nil,fmt.Errorf("Error getting template %s: %q",templFile,err)
	}
	sTpl := strings.ReplaceAll(hnd.Base,"{{.BODY}}",string(body))
	var tpl = template.New(name)
	tpl,err = tpl.Parse(string(sTpl))
	if err != nil{
		return nil,fmt.Errorf("Error parsing body to template: %q",err)
	}
	return tpl,nil
}
