package controllers

import (
	"net/http"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
)

func RenderMarkdown(p chan stuff.Document, publicationChan chan stuff.Document) {
	// lo que hace es parsear el markdown en html
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	// for now im doing this
	// but i want to use this with a db

	d := <-publicationChan
	// ya sabe, concurrencia
	// obtiene el markdown
	d.Body = string(markdown.ToHTML([]byte(d.Body), parser, nil)) // despues lo pasa a html
	p <- d                                                        // al final hace lo siguiente

}
func RenderInfo(w http.ResponseWriter,id int){
	p := make(chan stuff.Document)
		// then decode the markdown to html
		d := make(chan stuff.Document)
		go dataControl.GetOnlyOnePublication(id, d)
		go RenderMarkdown(p, d)
		t, _ := template.ParseFiles("view/template.html")
		// the goroutines are the best
		//aqui estamos usando templates para evitar que tener que estar usando otra cosa
		t.Execute(w, <-p)
}

