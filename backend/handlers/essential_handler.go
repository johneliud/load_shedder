package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func EssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "This is a critical service and always available.")
}
