package database

import (
	"fmt"
	"github.com/RianIhsan/go-postgres-redis/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func PostgresConnection(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DbHost,
		config.DbUsername,
		config.DbPassword,
		config.DbName,
		config.DbPort,
		config.DbSSL,
		config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
		return nil
	}

	fmt.Println("Connecting database successfully")
	return db
}
