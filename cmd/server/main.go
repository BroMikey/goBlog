package main

import (
	"fmt"
	"os"

	"github.com/BroMikey/goBlog/bootstrap"
)

func main() {
	configPath := os.Getenv("APP_CONFIG")

	Conf, err := bootstrap.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(Conf)

	logger := bootstrap.InitLogger(Conf)
	logger.Info("Logger initialized")

	db := bootstrap.InitGorm(Conf, logger)
	if db != nil {
		logger.Info("GORM initialized")
	} else {
		logger.Warn("GORM not initialized due to missing MySQL configuration")
		panic("MySQL configuration is required to proceed")
	}
}
