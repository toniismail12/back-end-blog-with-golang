package database

import (
	"log"
	"os"
	"project-pertama/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error load .env file")
	}
	dsn:=os.Getenv("DSN")
	database,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal connet database")
	}else{
		log.Println("connect success")
	}
	DB=database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}