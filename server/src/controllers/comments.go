package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ranon-rat/silent-songs/src/dataControl"
	"github.com/ranon-rat/silent-songs/src/stuff"
	"golang.org/x/sync/errgroup"
)

func NewComment(w http.ResponseWriter,r *http.Request,idPubl int)error {
	var comment stuff.Comment
	var controlErrors errgroup.Group
	json.NewDecoder(r.Body).Decode(&comment)
	controlErrors.Go(func() error {
		return dataControl.AddComment(comment, idPubl)
	})
	if err := controlErrors.Wait(); err != nil {
		
		return errors.New(fmt.Sprint(err.Error(), r.Header.Get("x-forwarded-for")))
	}
	w.Write([]byte("success upload"))
	return nil
}
