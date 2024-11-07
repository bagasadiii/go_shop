package config

import (
	"auth_service/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pass = "postgres"
	name = "auth_service"
)
func InitDBTest()*gorm.DB{
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,pass,name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return DB
}
