package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
	"golang.org/x/sync/errgroup"
)

// this is the post manager , with this you can do really interesting things
func NewPost(w http.ResponseWriter, r *http.Request) {

	// aqui solo es para ver los metodos
	canPublic := make(chan bool)
	cookie, err := r.Cookie("ip")
	if err != nil {
		cookie = &http.Cookie{
			Name:    "ip",
			Value:   r.Header.Get("x-forwarded-for"),
			Expires: time.Now().AddDate(1, 0, 0),
		}
		http.SetCookie(w, cookie)

	}

	dataControl.VerifyCookieIP(cookie.Value, canPublic)
	// new things
	// go back to the normal
	if <-canPublic {
		switch r.Method {
		case "POST":
			// i need to do some data bases for do this

			// in the future i gonna do something with this
			d, controlErrors, cont := stuff.Document{}, errgroup.Group{}, make(chan bool)

			// decode the bodyRequest
			json.NewDecoder(r.Body).Decode(&d)

			// this is for check if something is wrong
			go Check(cont, d, w)
			if <-cont {
				log.Println("fuck")
				return
			}
			controlErrors.Go(func() error {
				// add the publications
				return dataControl.AddPublication(d)
			})
			if err = controlErrors.Wait(); err != nil {
				log.Println(err)
				return
			}
			return

		case "GET":

			http.ServeFile(w, r, "view/post.html")
			return

		default:
			// solo acepta 2 metodos de request
			w.Write([]byte("ñao ñao voce es maricon"))
			break
		}
		return
	}
	//log.Println(conf)
	http.ServeFile(w, r, "view/denegado.html")

}
