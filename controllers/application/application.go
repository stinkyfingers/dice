package application

import (
	"github.com/stinkyfingers/dice/models/user"
	"html/template"
	// "log"
	"net/http"
	"strconv"
	"strings"
)

func Application(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	var userId int
	var u user.User

	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		userId = 0
	} else {
		userId, err = strconv.Atoi(strings.Split(cookie.String(), "=")[1])
	}
	if userId > 0 {
		u.ID = userId
		err = u.Get()
		if err != nil {
			http.Error(rw, "Error executing templates.", 400)
		}
		data["user"] = u
	}

	tname := "main"
	t, err := template.New(tname).ParseFiles("templates/main.tmpl", "templates/index.tmpl")
	err = t.ExecuteTemplate(rw, tname, data)
	if err != nil {
		http.Error(rw, "Error executing templates.", 400)
	}
}

func Login(rw http.ResponseWriter, r *http.Request) {
	tname := "main"
	t, err := template.New(tname).ParseFiles("templates/main.tmpl", "templates/login.tmpl")
	err = t.ExecuteTemplate(rw, tname, nil)
	if err != nil {
		http.Error(rw, "Error executing templates.", 400)
	}
}
