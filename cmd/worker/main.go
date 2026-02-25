package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/your-username/url-shortener/config"
	"github.com/your-username/url-shortener/internal/repository"
)

type ClickEvent struct {
	Code      string `json:"code"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

func main() {

	cfg := config.LoadConfig()
	mysqlRepo, err := repository.NewMSQLRepository(cfg.DBDSN)
	if err != nil {
		log.Fatalf("Worker Critical: MySQL connection failed: %v", err)
	}


	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   "url_clicks",
		GroupID: "analytics-group", 
	})

	log.Println(" Analytics Worker is running... ")

	
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Worker Error reading message: %v", err)
			continue
		}

		var event ClickEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Worker Error parsing JSON: %v", err)
			continue
		}


		err = mysqlRepo.SaveClick(context.Background(), event.Code, event.IP, event.UserAgent)
		if err != nil {
			log.Printf("Worker Error saving to DB: %v", err)
		} else {
			log.Printf("📊 Logged click: %s from %s", event.Code, event.IP)
		}
	}
}