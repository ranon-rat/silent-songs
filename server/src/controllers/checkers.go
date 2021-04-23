package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
	"golang.org/x/sync/errgroup"
)

// this only is for the styles and script
func Check(c chan bool, d stuff.Document, w http.ResponseWriter) {
	_, err := http.Get(d.Mineatura)
	log.Println(d, err)
	c <- d.Body == "" || d.Title == "" || d.Mineatura == "" || err != nil
}
func PostChecker(w http.ResponseWriter, r *http.Request)stuff.ErrorCode {
	// i need to do some data bases for do this

	// in the future i gonna do something with this
	d, controlErrors, cont := stuff.Document{}, errgroup.Group{}, make(chan bool)

	// decode the bodyRequest
	json.NewDecoder(r.Body).Decode(&d)

	// this is for check if something is wrong
	go Check(cont, d, w)
	if <-cont {
		log.Println("fuck")
		
		return stuff.ErrorCode{
			Code :400,
			Message: "something is invalid , try again",
		}
	}
	controlErrors.Go(func() error {
		// add the publications
		return dataControl.AddPublication(d)
	})
	if err := controlErrors.Wait(); err != nil {
	
		return stuff.ErrorCode{
			Code :500,
			Message: "Internal server error , contact with the admins",
		}
	}
	return stuff.ErrorCode{}

}