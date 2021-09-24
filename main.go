package main

import (
	"log"

	"github.com/BraulioAguilarDev/live-reloading-code/config"
)

func main() {

	cfg := config.Config

	if cfg.Debug {
		log.Printf("[BEDUG] Starting app: %s", cfg.AppName)
	}

	log.Println("Hello BraulioAguilarDev :)")
}
