package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3306)/grape_vs_martini_api")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Use(func(res http.ResponseWriter, req *http.Request, r render.Render) {
		api_key := ""
		api_key = req.URL.Query().Get("key")
		if api_key == "" {
			r.JSON(404, map[string]interface{}{"status": "Fail", "error_message": "Need api key"})
		} else {
			// r.JSON(200, map[string]interface{}{"key": api_key})
			current_company, company_id := GetCompany(db, api_key)
			if company_id < 0 {
				r.JSON(404, map[string]interface{}{"status": "Fail", "error_message": "Bad api key"})
			} else {
				m.Map(current_company)
			}
		}
	})

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/projects", func(current_company Company, r render.Render) {
		projects := GetProjects(db, current_company.Id)
		r.JSON(200, map[string]interface{}{"status": "Success", "data": projects})
	})

	m.Get("/projects/:id", func(current_company Company, params martini.Params, r render.Render) {
		paramId, err := strconv.Atoi(params["id"])
		if err != nil {
			r.JSON(404, map[string]interface{}{"status": "Fail", "error_message": err.Error()})
			return
		}
		project, id := GetProject(db, current_company.Id, paramId)
		if id > 0 {
			r.JSON(200, map[string]interface{}{"status": "Success", "data": project})
		} else {
			r.JSON(404, map[string]interface{}{"status": "Fail", "error_message": "Project not found"})
		}
	})

	// m.Run()
	http.ListenAndServe(":8080", m)
}
