package main

import (
	"fmt"
	"github.com/querybuilder/config"
	"github.com/querybuilder/database"
)

func main() {
	fmt.Println("Hello, World!")
    
    config := config.Get()
	database.Get(config.GetDBConnStr())
	database.Read()
}
