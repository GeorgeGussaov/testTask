package main

import (
	"log"
	"task/handlers"
	"task/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatalf("error initialization db: %v", err)
	}
	defer db.Close()

	subscriptionsRepo := repository.NewSubscriptionsRepository(db)
	r := gin.Default()

	r.Use(cors.New(cors.Config{ //гпт помог решить проблему со сваггером
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	handler := handlers.NewHandler(subscriptionsRepo)
	handler.RegisterRoutes(r)

	r.Run(":8080") //в докере пробрасывается на 8081
}
