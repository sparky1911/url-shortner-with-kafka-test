package config

import "os"

type Config struct {
	ServerPort  string
	DBDSN       string
	RedisAddr   string
	KafkaBroker string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DBDSN:       getEnv("DB_DSN", "root:pass@tcp(localhost:3306)/url_shortener?parseTime=true"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
