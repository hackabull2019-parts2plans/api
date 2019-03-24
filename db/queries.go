package Database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../models"
	"os"
	"errors"
)

var db *sql.DB

func Init() {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")

	if user == "" || dbname == "" {
		fmt.Println("[ERROR] Must provide username and database name!")
		os.Exit(-1)
	}

	conn, err := sql.Open("mysql", user + ":" + pass + "@/" + dbname)
	if err != nil {
		fmt.Println("[ERROR] Failed to connect to database!")
		panic(err.Error())
	}
	db = conn
}

func Close() {
	err := db.Close()

	if err != nil {
		fmt.Println("[ERROR] Failed to close database connection!")
		panic(err.Error())
	}
}

// var GetAllProjects = func() []models.Projects


func GetAllParts() []*models.Part {
	var items = []*models.Part{}

	rows, err := db.Query("SELECT partID, partName, partDesc FROM Part;")
	if err != nil {
		fmt.Println("[ERROR] GetAllParts query failed to execute")
		panic(err.Error())
	}

	for rows.Next() {
		var p models.Part

		err = rows.Scan(&p.Id, &p.Name, &p.Desc)
		if err != nil {
			fmt.Println("[ERROR] Failed to scan row into Part")
			panic(err.Error())
		}
		items = append(items, &p)
	}

	return items
}

func InsertPart(p models.Part) error {
	if p.Name == "" {
		fmt.Println("[WARN] Attempted to insert part with empty name")
		return errors.New("Tried to insert part with no name")
	}

	if p.Desc == "" {
		fmt.Println("[WARN] Attempted to insert part with no description")
		return errors.New("Tried to insert a part with no description")
	}

	_, err := db.Query("INSERT INTO Part (partName, partDesc) VALUES (?, ?);", p.Name, p.Desc);
	if err != nil {
		fmt.Println("[WARN] Failed to insert part into database")
		return errors.New("Failed to insert part into database")
	}
	fmt.Println("[LOG] Inserted part into database with data: name: " + p.Name + ", desc: " + p.Desc)
	return nil
}

func GetAllProjects() []*models.Project {
	var projects = []*models.Project{}

	rows, err := db.Query("SELECT projectID, projectName, projectDesc, imgPath, url FROM Project;")
	if err != nil {
		fmt.Println("[ERROR] GetAllProjects query failed to execute")
		panic(err.Error())
	}

	for rows.Next() {
		var p = models.Project{}

		err = rows.Scan(&p.Id, &p.Name, &p.Desc, &p.ImagePath, &p.Url)
		if err != nil {
			fmt.Println("[ERROR] Failed to scan row into Project")
			panic(err.Error())
		}
		projects = append(projects, &p)
	}

	return projects
}
