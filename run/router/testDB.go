package router

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	"log"
)

func TestingDb() {
	_, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, "artem", err)
		return
	}
	log.Println("успешно")
}
