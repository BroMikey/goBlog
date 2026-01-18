package main

import (
	"os"

	"github.com/BroMikey/goBlog/bootstrap"
)

func main() {
	configPath := os.Getenv("APP_CONFIG")

	Conf, err := bootstrap.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	logger := bootstrap.InitLogger(Conf)
	logger.Info("Logger initialized")

	db := bootstrap.InitGorm(Conf, logger)
	if db != nil {
		logger.Info("GORM initialized")
	} else {
		logger.Warn("GORM not initialized due to missing MySQL configuration")
		panic("MySQL configuration is required to proceed")
	}

	router := bootstrap.InitRouter(Conf, db, logger)
	logger.Info("Router initialized")

	addr := Conf.System.Addr()
	logger.Infof("Starting server at %s", addr)
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
