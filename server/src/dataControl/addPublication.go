package dataControl

import (
	"errors"

	"github.com/ranon-rat/silent-songs/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)


func AddPublication(e stuff.Document) error {
	q := `INSERT INTO 
	publ(title,mineatura,body) 
	values(?1,?2,?3);
	`
	db := GetConnection()
	defer db.Close()
	stm, _ := db.Prepare(q)
	
	defer stm.Close()
	r, _ := stm.Exec(&e.Title, &e.Mineatura, &e.Body)
	
	i, _ := r.RowsAffected()
	if i != 1 {
		
		return errors.New("se esperaba una sola fila omg")
	}
	return nil
}
