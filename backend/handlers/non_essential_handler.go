package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func NonEssentialHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "Non-essential service responded successfully!\nThis service is available when load is manageable.")
}
