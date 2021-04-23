package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Publication(w http.ResponseWriter, r *http.Request) {
	attr := mux.Vars(r)

	// get the id of the publication
	idPubl, _ := strconv.Atoi(attr["id"])
	switch r.Method {
	case "GET":
		RenderInfo(w,idPubl)
		return
	case "POST":
		if err:=NewComment(w,r,idPubl);err!=nil{
			log.Println(err)
			return
		}
		return
	}
	w.Write([]byte("ñao ñao voce e maricon"))
}