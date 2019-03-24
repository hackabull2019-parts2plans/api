package controllers

import (
	"fmt"
	"net/http"
	// db "../db"
)

var findProjects = func(w http.ResponseWriter, r *http.Request) {
	
	// projects := db.GetAllProjects()

	fmt.Fprintf(w, "Test response")
}
