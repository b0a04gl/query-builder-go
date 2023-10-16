package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	dbUser     string
	dbName     string
}

func Get() *Config {

    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	postgresUser := os.Getenv("POSTGRES_USER")

	fmt.Println("user : "+postgresUser)

	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB name")
	
	flag.Parse()

	return conf
}


func (c *Config) GetDBConnStr() string {
	return fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable",
		c.dbUser,
		c.dbName,
	)
}
