package handlers

import (
	"net/http"
	"time"
)

func SimulateLoadHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate some processing time to make the load simulation more realistic
	time.Sleep(50 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
}
