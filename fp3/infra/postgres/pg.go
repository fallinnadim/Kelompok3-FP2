package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/adenhidayatuloh/Kelompok5_FinalProject3/entity"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	db       *gorm.DB
	err      error
)

func GetDBConfig() gorm.Dialector {

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		port,
		user,
		password,
		dbname,
	)

	return postgres.Open(dbConfig)
}

func GetDBInstance() *gorm.DB {
	return db
}

func seedAdmin() {
	admin := &entity.User{
		FullName: "admin",
		Email:    "admin@gmail.com",
		Password: "admin12345",
		Role:     "admin",
	}
	err := admin.HashPassword()

	if err != nil {
		fmt.Println(err)
	}

	err2 := db.Create(admin).Error

	if err2 != nil {
		fmt.Println(err2)
	}

	log.Println("Admin account seed success!")
}

func init() {
	db, err = gorm.Open(GetDBConfig(), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{})

	if db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedAdmin()
		}
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	log.Println("Connected to DB!")
}
