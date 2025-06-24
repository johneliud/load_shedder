package handlers

import (
	"net/http"
	"time"
)

func SimulateLoadHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate longer processing time to maintain load for demo purposes
	time.Sleep(2 * time.Second)
	w.WriteHeader(http.StatusOK)
}
