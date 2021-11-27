package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
)

var db *gorm.DB

func init() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, username, dbName, password))
	if err != nil {
		log.Fatal(err)
	} else {
		db = conn
		log.Info("DB connected on " + dbHost)
	}

	err = migrateSchema()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Schema migrated successfully")
	}
}

func GetDB() *gorm.DB {
	return db
}

func migrateSchema() error {
	// Notice: many-to-many first
	err := db.AutoMigrate(

	).Error
	return err
}
