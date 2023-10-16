package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/querybuilder/model"
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

func Read(query string, args any) ([]model.Player){

	if GlobalDB == nil {
		log.Fatal("DB not available")
	}

	rows, err := GlobalDB.Client.Query(query, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var players []model.Player

	for rows.Next() {
		var player model.Player
		if err := rows.Scan(&player.Id, &player.Name, &player.Team, &player.Role, &player.Total_runs, &player.Total_wickets); err != nil {
			log.Fatal(err)
		}
		players = append(players, player)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return players
}
