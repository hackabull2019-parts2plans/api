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

func InsertProject(p models.Project) (int64, error) {
	if p.Name == "" {
		fmt.Println("[WARN] Attempted to insert project with empty name")
		return 0, errors.New("Tried to insert project with no name")
	}

	if p.Desc == "" {
		fmt.Println("[WARN] Attempted to insert project with no description")
		return 0, errors.New("Tried to insert a project with no description")
	}

	res, err := db.Exec("INSERT INTO Project (projectName, projectDesc, imgPath, url) VALUES (?, ?, ?, ?)", p.Name, p.Desc, p.ImagePath, p.Url);
	if err != nil {
		fmt.Println("[WARN] Failed to insert project into database")
		return 0, errors.New("Failed to insert project into database")
	}
	fmt.Println("[LOG] Inserted Project into database with data: name: " + p.Name + ", desc: " + p.Desc)
	id, _ := res.LastInsertId()
	return id, nil
}

func getLastID() int {
	var id int
	err := db.QueryRow("SELECT LAST_INSERT_ID();").Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return id
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

	for _, project := range projects {
		project.Parts,_ = GetParts(project.Id)
	}

	return projects
}

func GetParts(projectId int) ([]*models.Part, error) {
	var parts = []*models.Part{}
	rows, err := db.Query("SELECT Part.partID, Part.partName, CompMapping.qty FROM Part LEFT JOIN CompMapping ON Part.partID = CompMapping.partID WHERE projectID = ?", projectId)
	if err != nil {
		fmt.Println("[ERROR] GetParts query failed to execute")
		panic(err.Error())
	}

	for rows.Next() {
		var p models.Part

		err = rows.Scan(&p.Id, &p.Name, &p.Qty)
		if err != nil {
			fmt.Println("[ERROR] Failed to scan row into Part")
			panic(err.Error())
		}
		parts = append(parts, &p)
	}
	return parts, nil
}

func AddPart(project models.Project, part *models.Part) error {
	if project.Id == 0 {
		return errors.New("Invalid project id: 0")
	}

	if part.Id == 0 {
		return errors.New("Invalid part id: 0")
	}

	if part.Qty == 0 {
		return errors.New("Invalid quantity of parts")
	}

	_, err := db.Query("INSERT INTO CompMapping (projectID, partID, qty) VALUES (?, ?, ?);", project.Id, part.Id, part.Qty);
	if err != nil {
		fmt.Println("[WARN] Failed to insert part to project mapping into database")
		return errors.New("Failed to insert part to project mapping into database")
	}

	return nil
}
