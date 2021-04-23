package dataControl

import (
	"fmt"

	"github.com/ranon-rat/silent-songs/src/stuff"
)

func AddComment(comment stuff.Comment, IDPubl int) error {
	q := `INSERT INTO comment(id_publ,username,body)
	VALUES(?1,?2,?3)
	`
	db := GetConnection()
	defer db.Close()
	stm, _ := db.Prepare(q)
	defer stm.Close()
	rows, _ := stm.Exec(IDPubl, comment.Username, comment.Body)
	if howMany, _ := rows.RowsAffected(); howMany > 1 {
		return fmt.Errorf("Username:%s\nBody:%s", comment.Username, comment.Body)

	}
	return nil

}
