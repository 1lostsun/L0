package app

import (
	"ProducerService/internal/handler"
	"ProducerService/internal/kafka"
	"ProducerService/internal/router"
	"ProducerService/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func Run() {
	if err := godotenv.Load("ProducerService/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	engine := gin.Default()

	kafkaCfg := kafka.NewKafkaCfg()
	uc := usecase.New(kafkaCfg)
	h := handler.NewHandler(uc)
	r := router.NewRouter(engine, h)
	r.InitRoutes()
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
