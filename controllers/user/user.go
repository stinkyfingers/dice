package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/stinkyfingers/dice/models/user_mgo"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
)

const (
	timeFormat = "Jan 2, 2006 at 3:04pm (MST)"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	var u user_mgo.User
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = u.Create()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	var u user_mgo.User
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = u.Get()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var u user_mgo.User
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = u.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func AuthenticateUser(rw http.ResponseWriter, r *http.Request) {
	var u user_mgo.User
	var cookie http.Cookie
	var err error
	ct := r.Header.Get("Content-Type")
	if ct == "application/json" {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), 404)
		}
		err = json.Unmarshal(requestBody, &u)
		if err != nil {
			http.Error(rw, err.Error(), 404)
		}
	} else {
		u.Email = r.FormValue("email")
		u.Password = r.FormValue("password")
	}
	err = u.Authenticate()
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("Invalid email or password")
		}
		http.Error(rw, err.Error(), 404)
	}

	cookie.Name = "user"
	cookie.Value = u.ObjectID.Hex()
	cookie.Expires = time.Now().AddDate(0, 0, 1)
	http.SetCookie(rw, &cookie)

	http.Redirect(rw, r, "/", 301)
}

func Logout(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	cookie.MaxAge = -1
	http.SetCookie(rw, cookie)
	http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
}

func Register(rw http.ResponseWriter, r *http.Request) {
	var err error
	var u user_mgo.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	err = u.Create()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func ResetPassword(rw http.ResponseWriter, r *http.Request) {
	var err error
	var u user_mgo.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = u.ResetPassword()
	if err != nil {
		return

	}

	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	rw.Write(jstring)

}
