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

	fmt.Printf("connected to %v\n\n", GlobalDB.Client)
}

func Read(query string,args any) {

	if GlobalDB == nil{
		log.Fatal("DB not available")
	}

	rows, err := GlobalDB.Client.Query(query,args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var team string
		var role string
		var total_runs int
		var total_wickets int
		if err := rows.Scan(&id, &name, &team, &role, &total_runs, &total_wickets); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\nID: %d, Name: %s, team: %s, role: %s, total_runs: %d, total_wickets : %d \n", id, name, team, role, total_runs, total_wickets)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
