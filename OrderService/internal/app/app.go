package app

import (
	"OrderService/OrderService/internal/config"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func Run() {
	if err := godotenv.Load("OrderService/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.KafkaCfg.Topic)

	//db, err := postgres.NewDBConfig(cfg.PgCfg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//kafkaCfg := kafka.NewKafkaCfg(cfg.KafkaCfg)
}
