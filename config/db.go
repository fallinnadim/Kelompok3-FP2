package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func StartDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to database")
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		age SMALLSERIAL NOT NULL,
		created_at DATE NOT NULL,
		updated_at DATE NOT NULL
	);
	CREATE TABLE IF NOT EXISTS social_medias (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		social_media_url VARCHAR(50) NOT NULL,
		user_id SERIAL NOT NULL,
		CONSTRAINT fk_user_id
  		FOREIGN KEY(user_id)
  		REFERENCES users(id)
	);
	`
	_, errQuery := db.Exec(string(query))
	if errQuery != nil {
		log.Panicln("Failed to Create table")
	}
	return db
}
