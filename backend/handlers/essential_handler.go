package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func EssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	
	// Get simulated load from query parameter
	simulatedLoadStr := r.URL.Query().Get("simulatedLoad")
	simulatedLoad := 0
	if simulatedLoadStr != "" {
		if load, err := strconv.Atoi(simulatedLoadStr); err == nil {
			simulatedLoad = load
		}
	}
	
	fmt.Fprintf(w, "Essential service responded successfully! This critical service is always available, even under high load (%v) requests.\n", simulatedLoad)
}
