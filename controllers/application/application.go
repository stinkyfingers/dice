package application

import (
	// "encoding/json"
	// "github.com/stinkyfingers/dice/models/user"
	// "io/ioutil"
	// "log"
	"html/template"
	"net/http"
)

func Application(rw http.ResponseWriter, r *http.Request) {
	tname := "main"
	t, err := template.New(tname).ParseFiles("templates/main.tmpl", "templates/index.tmpl")
	err = t.ExecuteTemplate(rw, tname, nil)
	if err != nil {
		http.Error(rw, "Error executing templates.", 404)
	}
}
