package main

import (
	"fmt"

	"github.com/querybuilder/config"
	"github.com/querybuilder/database"
	"github.com/querybuilder/querybuilder"
)

func main() {
	fmt.Println("Hello, World!")
    
    config := config.Get()
	database.Get(config.GetDBConnStr())
    
	builder := querybuilder.NewSQLBuilder()

	query, args := builder.
		Select("id", "name", "team", "role" , "total_runs", "total_wickets").
		From("players").
		Where("role = $1","batsman").
		OrderBy("total_runs DESC").
		Build()

	fmt.Println("\nGenerated SQL query:", query)
	fmt.Println("Query arguments:", args[0])

	database.Read(query,args[0])
}
