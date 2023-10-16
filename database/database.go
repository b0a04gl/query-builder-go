package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	Client *sql.DB
}

var GlobalDB *DB

func Get(connStr string) {

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return
	}

	if err := db.Ping(); err != nil {
		return
	}

	GlobalDB = &DB{
		Client: db,
	}

	fmt.Printf("connected to %v\n", GlobalDB.Client)
}

func Read() {

	if GlobalDB == nil{
		log.Fatal("DB not available")
	}

	rows, err := GlobalDB.Client.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
