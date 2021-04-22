package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/silent-songs/src/controllers"
)

func Routes() error {
	// aqui solo es para dar la salida de informacion
	r := mux.NewRouter()
	// REDIRIJE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/1", 301)
		// yeah this is cool
	})
	// load the page
	r.HandleFunc(`/{page:[\d]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/home.html")
	})
	// send the files
	r.HandleFunc(`/public/{file:[\/\w\W]+?}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	r.HandleFunc("/admin/postfile", controllers.NewPost)
	// get the info
	r.HandleFunc(`/api/{page:[\d]+}`, controllers.Api)
	// render the publication

	r.HandleFunc(`/publication/{id:[\d]+}`, controllers.RenderInfo)
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	return http.ListenAndServe(":"+port, r)
}
