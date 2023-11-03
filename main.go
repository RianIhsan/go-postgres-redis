package main

import (
	"fmt"
	"github.com/RianIhsan/go-postgres-redis/config"
	"github.com/RianIhsan/go-postgres-redis/database"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}
	db := database.PostgresConnection(&loadConfig)
	rdb := database.RedisConnection(&loadConfig)
	bootApp(db, rdb)
}

func bootApp(db *gorm.DB, rdb *redis.Client) {
	app := gin.Default()
	err := app.Run()
	if err != nil {
		panic("Cannot run app")
	}
	fmt.Println("Server started")
}
