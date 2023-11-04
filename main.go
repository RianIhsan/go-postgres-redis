package main

import (
	"fmt"
	"github.com/RianIhsan/go-postgres-redis/config"
	"github.com/RianIhsan/go-postgres-redis/controller"
	"github.com/RianIhsan/go-postgres-redis/database"
	"github.com/RianIhsan/go-postgres-redis/model"
	"github.com/RianIhsan/go-postgres-redis/repository"
	"github.com/RianIhsan/go-postgres-redis/router"
	"github.com/RianIhsan/go-postgres-redis/usecase"
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
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatal("Failed to migrate")
	}
	rdb := database.RedisConnection(&loadConfig)
	bootApp(db, rdb)
}

func bootApp(db *gorm.DB, rdb *redis.Client) {
	app := gin.Default()

	bookRepo := repository.NewBookRepository(db, rdb)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookController := controller.NewBookController(bookUsecase)

	route := router.NewRouter(app, bookController)
	err := route.Run()
	if err != nil {
		panic("Cannot run app")
	}
	fmt.Println("Server started")
}
