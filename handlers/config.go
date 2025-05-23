package handlers

import (
	"encoding/json"
	"net/http"
	"your-module-name/config"
)

var AppConfig *config.Config

func SetConfig(c *config.Config) {
	AppConfig = c
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AppConfig)
}
