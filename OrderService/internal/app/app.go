package app

import (
	"OrderService/OrderService/internal/config"
	"OrderService/OrderService/internal/handler"
	"OrderService/OrderService/internal/kafka"
	"OrderService/OrderService/internal/repository/postgres"
	"OrderService/OrderService/internal/repository/redis"
	"OrderService/OrderService/internal/router"
	"OrderService/OrderService/internal/service"
	"OrderService/OrderService/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	if err := godotenv.Load("OrderService/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewDBConfig(cfg.PgCfg)
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	kafkaCfg := kafka.NewKafkaCfg(cfg.KafkaCfg)
	pgRepository := postgres.NewOrderRepository(db)
	redisRepository := redis.NewConfig(cfg.RedisCfg)
	orderService := service.New(pgRepository, redisRepository)
	orderUsecase := usecase.New(kafkaCfg, orderService)
	orderHandler := handler.NewOrderHandler(orderService)
	orderRouter := router.NewRouter(engine, orderHandler)
	orderRouter.InitRoutes()

	go func() {
		err := orderUsecase.ReadKafkaMessage(ctx)
		if err != nil {
			log.Println("kafka consumer stopped with err: ", err)
		}
	}()

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- engine.Run(":8081")
	}()

	select {
	case <-sigs:
		log.Println("Shutting down server")
		cancel()
	case err := <-serverErr:
		log.Fatalf("HTTP server error: %v", err)
	}
}
