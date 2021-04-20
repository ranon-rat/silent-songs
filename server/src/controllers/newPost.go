package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ranon-rat/silent-songs/src/dataControll"
	"github.com/ranon-rat/silent-songs/src/stuff"
)

// this is the post manager , with this you can do really interesting things
func NewPost(w http.ResponseWriter, r *http.Request) {
	canPublic:=make(chan bool)

	// aqui solo es para ver los metodos

	cookie, err := r.Cookie("ip")
	if err != nil {

		http.SetCookie(w, &http.Cookie{
			Name:    "ip",
			Value:   r.Header.Get("x-forwarded-for"),
			Expires: time.Now().AddDate(1, 0, 0),
		})
	}
	go dataControll.VerifyCookieIP(cookie.Value,canPublic)
	// new things
	// go back to the normal
	if <-canPublic {
		switch r.Method {
		case "POST":
			// i need to do some data bases for do this

			// in the future i gonna do something with this
			var d stuff.Document

			// decode the bodyRequest
			json.NewDecoder(r.Body).Decode(&d)
			cont := make(chan bool)
			// this is for check if something is wrong
			go Check(cont, d, w)
			if <-cont {
				log.Println("fuck")
				return
			}
			// add the publications
			go dataControll.AddPublication(d)
			break
		case "GET":

			http.ServeFile(w, r, "view/post.html")
			break
		default:
			// solo acepta 2 metodos de request
			w.Write([]byte("ñao ñao voce es maricon"))
			break
		}
	} else {

		http.ServeFile(w, r, "view/denegado.html")
	}

}