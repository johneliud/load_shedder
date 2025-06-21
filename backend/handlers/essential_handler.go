package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func EssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "Essential service responded successfully!\nThis critical service is always available, even under high load.")
}
