package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	
	err = godotenv.Load(".env") 
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host		= "localhost"
		port		= 5432
		user 		= os.Getenv("DB_USER")
		password	= os.Getenv("DB_PASSWORD")
		dbname		= os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database.")
}

func GetDB() *sql.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}