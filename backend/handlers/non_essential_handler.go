package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func NonEssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "This is a non-essential service and may be dropped.")
}
