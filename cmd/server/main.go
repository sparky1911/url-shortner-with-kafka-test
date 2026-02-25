package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/your-username/url-shortener/config"
	"github.com/your-username/url-shortener/internal/api"
	"github.com/your-username/url-shortener/internal/api/handler"
	"github.com/your-username/url-shortener/internal/repository"
	"github.com/your-username/url-shortener/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	mysqlRepo, err := repository.NewMSQLRepository(cfg.DBDSN)
	if err != nil {
		log.Fatalf("Critical: MySQL connection failed: %v", err)
	}

	redisRepo := repository.NewRedisRepository(cfg.RedisAddr)
	kafkaRepo:=repository.NewKafkaRepository(cfg.KafkaBroker)

	urlSvc := service.NewURLService(mysqlRepo, redisRepo,*kafkaRepo)
	urlHandler := handler.NewURLHandler(urlSvc)

	r := gin.Default()

	r = api.SetupRouter(urlHandler)

	log.Printf("Application initialized. Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Critical: Server failed to start: %v", err)
	}
}


