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

}
