package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func NonEssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	
	// Get simulated load from query parameter
	simulatedLoadStr := r.URL.Query().Get("simulatedLoad")
	simulatedLoad := 0
	if simulatedLoadStr != "" {
		if load, err := strconv.Atoi(simulatedLoadStr); err == nil {
			simulatedLoad = load
		}
	}
	
	if simulatedLoad > 50 {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "Non-essential service unavailable! Load shedding active - service temporarily unavailable due to high load.")
		return
	}
	
	fmt.Fprintln(w, "Non-essential service responded successfully! This service is available when load is manageable.")
}
