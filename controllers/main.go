package controllers

import (
	"net/http"
	"encoding/json"
	db "../db"
)

type Route struct{
	Path	string
	Method	string
	Handler	http.HandlerFunc
}

var Routes = []*Route{
	&Route{
		Path: "/findProjects",
		Method: "GET",
		Handler: findProjects,
	},
	&Route{
		Path: "/parts",
		Method: "GET",
		Handler: getAllParts,
	},
}

func getAllParts(w http.ResponseWriter, r *http.Request) {
	parts := db.GetAllParts()

	json.NewEncoder(w).Encode(parts)
}
