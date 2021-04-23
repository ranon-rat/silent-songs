package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
)

// this is the ap

func Api(w http.ResponseWriter, r *http.Request) {
	// only send this
	// this is for use the apis
	min, _ := strconv.Atoi(mux.Vars(r)["page"])

	// concurrency communication
	//the db management
	sizeChan, dChan := make(chan int), make(chan []stuff.Document)

	// we use this function only one time so, im only usign a anon function 😩

	go dataControl.GetPublications(min, dChan)

	go dataControl.GetTheSizeOfTheQuery(sizeChan)

	api := stuff.Publications{

		Quantity: stuff.Quantity,
		Publications: <-dChan,
		Size:         <-sizeChan,
	}
	// send the json
	json.NewEncoder(w).Encode(api)

}

// hello there
// this is only the setup
