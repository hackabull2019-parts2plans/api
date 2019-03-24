package controllers

import (
	"net/http"
	"encoding/json"
	db "../db"
	"fmt"
	"io/ioutil"
	"../models"
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
	&Route{
		Path: "/projects",
		Method: "GET",
		Handler: getAllProjects,
	},
	&Route{
		Path: "/part",
		Method: "POST",
		Handler: insertPart,
	},
	&Route{
		Path: "/project",
		Method: "POST",
		Handler: addProject,
	},
}

func lookupProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func getAllParts(w http.ResponseWriter, r *http.Request) {
	parts := db.GetAllParts()

	json.NewEncoder(w).Encode(parts)
}

func getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := db.GetAllProjects()

	json.NewEncoder(w).Encode(projects)
}

func insertPart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("[LOG] Failed to read body of message")
		w.WriteHeader(400)
		return
	}

	var p models.Part

	err = json.Unmarshal(body, &p)
	if err != nil {
		// Invalid json
		fmt.Println("[LOG] Invalid json")
		w.WriteHeader(400)
		return
	}

	err = db.InsertPart(p)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("[LOG] Failed to read body of message")
		w.WriteHeader(400)
		return
	}

	var p models.Project

	err = json.Unmarshal(body, &p)
	if err != nil {
		// Invalid json
		fmt.Println("[LOG] Invalid json")
		w.WriteHeader(400)
		return
	}

	pro, _ := json.Marshal(p)

	fmt.Println(string(pro))

	id, _ := db.InsertProject(p)

	p.Id = int(id)

	for _, part := range p.Parts {
		fmt.Println("Adding part")
		err = db.AddPart(p, part)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	w.WriteHeader(200)
}
