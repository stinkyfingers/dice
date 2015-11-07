package application

import (
	"github.com/stinkyfingers/dice/models/user_mgo"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"strings"
)

func Application(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	var userId bson.ObjectId
	var u user_mgo.User

	cookie, err := r.Cookie("user")
	data["user"] = nil
	if err == nil && cookie != nil {
		var uid string
		uid = strings.Split(cookie.String(), "=")[1]
		if err != nil {
			return
		}
		if uid != "" {
			userId = bson.ObjectIdHex(uid)
		}

		u.ObjectID = userId
		err = u.Get()
		if err != nil {
			http.Error(rw, "Error getting user.", 400)
		}

		data["user"] = u
	}
	tname := "main"
	t, err := template.New(tname).ParseFiles("templates/main.tmpl", "templates/index.tmpl")
	if err != nil {
		http.Error(rw, "Error parsing templates.", 400)
	}
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
func Logout(rw http.ResponseWriter, r *http.Request) {
	var cookie http.Cookie
	cookie.Name = "user"
	cookie.MaxAge = -1
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, r, "/", 301)
}
