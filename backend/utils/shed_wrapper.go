package utils

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var (
	maxRequests     = int32(50) // Maximum concurrent allowed requests
	currentRequests int32       // Global counter for active requests shared across all handlers
)

func ShedWrapper(h http.HandlerFunc, canShed bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Atomically check and increase current requests
		if atomic.AddInt32(&currentRequests, 1) > maxRequests {
			atomic.AddInt32(&currentRequests, -1)

			if canShed {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintln(w, "This service is temporarily unavailable due to high load.")
				return
			}
		}
		defer atomic.AddInt32(&currentRequests, -1)
		h(w, r)
	}
}
