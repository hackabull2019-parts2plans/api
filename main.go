package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"./controllers"
	db "./db"
)

func main() {

	db.Init()

	router := mux.NewRouter()

	for _, route := range controllers.Routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on port: " + port)

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	err := http.ListenAndServe(host + ":" + port, router)
	if err != nil {
		fmt.Print(err)
	}

	db.Close()
}
