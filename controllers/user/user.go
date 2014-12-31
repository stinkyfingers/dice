package user

import (
	"encoding/json"
	"github.com/stinkyfingers/dice/models/user"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	var u user.User
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = u.CreateUser()
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
	var u user.User
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
	var u user.User
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
	var u user.User
	var cookie http.Cookie
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = u.Authenticate()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	cookie.Name = "user"
	cookie.Value = u.ID
	cookie.Expires = 86400 * time.Second
	http.SetCookie(rw, cookie)

	jstring, err := json.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func Logout(rw http.Response, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	cookie.MaxAge = -1
	http.SetCookie(rw, cookie)
	rw.Write("success")
}
