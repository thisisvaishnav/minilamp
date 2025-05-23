package main

import (
	"fmt"
	"log"
	"net/http"
	"your-module-name/config"
	"your-module-name/handlers"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	handlers.SetConfig(cfg)

	http.HandleFunc("/config", handlers.GetConfigHandler)
	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
