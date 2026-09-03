package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing DB...")
	host, user, pass, dbName, port := GetDBConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		host, user, pass, dbName, port,
	)
	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic(errDb)
	}

	log.Println("DB initialized")

	DB = db
}

func GetDBConfig() (host, user, password, db, port string) {
	return viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"),
		viper.GetString("DB_PORT")
}
