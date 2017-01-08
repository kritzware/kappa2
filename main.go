package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "USER:PASSWORD@tcp(SERVER:3306)/DATABASE")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int
		username string
	)

	rows, err := db.Query("SELECT id, username FROM channels")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, username)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}