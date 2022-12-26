package views

import (
	"html/template"
	"log"
)

type IndexPage struct {
	Page *template.Template
	Name string
}

func (page IndexPage) RenderIndex() IndexPage {
	var err error
	page.Page, err = template.ParseFiles("./views/html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return page
}