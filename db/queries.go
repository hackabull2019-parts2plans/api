package Database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../models"
	"os"
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

