package dataControl

import (
	"github.com/ranon-rat/silent-songs/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)




func GetOnlyOnePublication(id int, aChan chan stuff.Document)  {
	q := (`SELECT * FROM publ WHERE id=?1`)
	db := GetConnection()
	defer db.Close()
	p, _ := db.Query(q, id)
	

	defer p.Close()
	var d stuff.Document
	for p.Next() {

		p.Scan(&d.ID, &d.Title, &d.Miniature, &d.Body)
		
	}
	aChan <- d
	
}