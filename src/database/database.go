package database

import (
	"fmt"
	"hello-gocd/src/models"
	"hello-gocd/src/utils"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var (
	Database DbInstance
	User     string = utils.GetDotEnv("DB_USER")
	Pass     string = utils.GetDotEnv("DB_PASS")
	Host     string = utils.GetDotEnv("DB_HOST")
	Port     string = utils.GetDotEnv("DB_PORT")
	Db       string = utils.GetDotEnv("DB_NAME")
)

func ConnectDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Pass, Host, Port, Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect do the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully!")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//TODO: add migrations

	db.AutoMigrate(&models.Names{})

	Database = DbInstance{Db: db}

}
