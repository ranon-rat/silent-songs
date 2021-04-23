package dataControl

import (
	_ "github.com/mattn/go-sqlite3"
)

func GetTheSizeOfTheQuery(sizeChan chan int) error {
	q := `
	SELECT COUNT(*) 
	FROM publ
	`
	// como no he encontrado muchas maneras de encontrar el
	//tamaño de una tabla lo que hace aqui es basicamente seleccionar el maximo valor

	db,dataSize  := GetConnection(),0
	defer db.Close()
	m, _ := db.Query(q)
	defer m.Close()
	for m.Next() {

		m.Scan(&dataSize)

	}
	sizeChan <- dataSize
	return nil
}
