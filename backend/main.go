package main

import (
	"fmt"
	"net/http"

	"github.com/johneliud/load_shedder/backend/handlers"
	"github.com/johneliud/load_shedder/backend/utils"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/essential", utils.ShedWrapper(handlers.EssentialHandler, false))
	http.HandleFunc("/non-essential", utils.ShedWrapper(handlers.NonEssentialHandler, true))

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
