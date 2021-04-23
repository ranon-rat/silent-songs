package controllers

import (
	"net/http"
	"time"

	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
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
		var void stuff.ErrorCode
		if err:=PostChecker(w , r );err!=void {
			http.Error(w,err.Message,err.Code)
			return
		}
		w.Write([]byte("succesfull upload"))

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

