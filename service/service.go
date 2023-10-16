package service

import (
	"fmt"

	"github.com/querybuilder/config"
	"github.com/querybuilder/database"
	"github.com/querybuilder/model"
	"github.com/querybuilder/querybuilder"
)

func Serve(request model.QueryRequest) (model.QueryResponse){
	config := config.Get()
	database.Get(config.GetDBConnStr())

	builder := querybuilder.NewSQLBuilder()

	query, args := builder.
		Select("id", "name", "team", "role", "total_runs", "total_wickets").
		From("players").
		Where("role = $1", "batsman").
		OrderBy("total_runs DESC").
		Build()

	fmt.Println("\nGenerated SQL query:", query)
	fmt.Println("Query arguments:", args)
	players := database.Read(query, args)
	fmt.Println(players)
	return model.QueryResponse{Players: players}
}