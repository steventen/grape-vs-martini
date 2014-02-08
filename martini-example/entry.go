package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Company struct {
	Id  int    `json:"id"`
	Api string `json:"api_key"`
}

func GetCompany(db *sql.DB, key string) (Company, int) {
	var company_id int
	var api_key string
	err := db.QueryRow("select id, api from companies where api = ? limit 1", key).Scan(&company_id, &api_key)
	switch {
	case err == sql.ErrNoRows:
		return Company{}, 0
	case err != nil:
		fmt.Println(err)
		return Company{}, -1
	default:
		return Company{company_id, api_key}, company_id
	}

}

func GetProject(db *sql.DB, company_id int, project_id int) (Project, int) {
	var (
		id   int
		name string
	)
	err := db.QueryRow("select id, name from projects where id = ? and company_id = ? limit 1", project_id, company_id).Scan(&id, &name)
	switch {
	case err == sql.ErrNoRows:
		return Project{}, 0
	case err != nil:
		fmt.Println(err)
		return Project{}, -1
	default:
		return Project{id, name}, id
	}

}

func GetProjects(db *sql.DB, companyId int) []Project {
	projects, err := db.Query("select id, name from projects where company_id = ?", companyId)
	if err != nil {
		fmt.Println(err)
	}

	var (
		id   int
		name string
	)

	p := make([]Project, 0)
	defer projects.Close()
	for projects.Next() {
		err := projects.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		} else {
			p = append(p, Project{id, name})
		}
	}
	return p
}
