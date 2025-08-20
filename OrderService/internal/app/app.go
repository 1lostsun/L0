package app

import (
	"OrderService/internal/config"
	"fmt"
	"log"
)

func Run() {
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
