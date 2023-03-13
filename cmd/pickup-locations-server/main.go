package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"

	"pavel-fokin/pickup-locations/internal/router"
	"pavel-fokin/pickup-locations/internal/server"
)

type Config struct {
	Server server.Config
}

func ReadConfig() *Config {
	envFile := "local.env"
	if os.Getenv("PICKUP_LOCATIONS_ENV_FILE") != "" {
		envFile = os.Getenv("PICKUP_LOCATIONS_ENV_FILE")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Printf("%v\n", err)
	}
	return cfg
}

// @Title Pickup Locations API
// @Version 0.0.1
// @Description Pickup Locations is a service that takes the source and a list of destinations and returns a list of routes between source and each destination.
// @BasePath /
func main() {
	config := ReadConfig()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	router := router.New()

	httpServer := server.New(config.Server)
	httpServer.SetupRoutesAPI(router)

	go httpServer.Start()

	<-sig

	httpServer.Shutdown()
}
