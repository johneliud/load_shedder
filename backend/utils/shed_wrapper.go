package utils

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func ShedWrapper(h http.HandlerFunc, canShed bool) http.HandlerFunc {
	var (
		maxRequests     = int32(10) // Maximum concurrent allowed requests
		currentRequests int32       // Counter for active requests
	)

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
