package dice

import (
	"encoding/json"
	"github.com/stinkyfingers/dice/models/dice_mgo"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"

	"log"
	"net/http"
	"strings"
)

func Roll(rw http.ResponseWriter, r *http.Request) {
	var ds dice_mgo.DiceSet
	var rs dice_mgo.Results

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		log.Print(err)
		http.Error(rw, err.Error(), 400)
	}

	rs, err = ds.Roll()

	jstring, err := json.Marshal(rs)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jstring)
}

func GetPublicDiceSets(rw http.ResponseWriter, r *http.Request) {
	var dss []dice_mgo.DiceSet
	var err error

	dss, err = dice_mgo.GetPublicDiceSets()
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	jstring, err := json.Marshal(dss)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	rw.Write(jstring)
}

func GetUserDiceSets(rw http.ResponseWriter, r *http.Request) {
	var dss []dice_mgo.DiceSet
	var userId bson.ObjectId
	var err error
	rw.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		rw.Write(nil)
		return
	}
	c := strings.Split(cookie.String(), "=")[1]

	if c != "" {
		userId = bson.ObjectIdHex(c)
	}

	dss, err = dice_mgo.GetUserDiceSets(userId) //was UserId
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	jstring, err := json.Marshal(dss)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	rw.Write(jstring)
}

func GetDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice_mgo.DiceSet

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		http.Error(rw, err.Error(), 400)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = ds.Get()
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	jstring, err := json.Marshal(ds)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jstring)
}

func SaveDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice_mgo.DiceSet

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		log.Print(err)
		http.Error(rw, err.Error(), 400)
	}

	if ds.ID.Valid() {
		err = ds.Update()
		if err != nil {
			log.Print("HERE", err)
			http.Error(rw, err.Error(), 400)
		}
	} else {
		err = ds.Create()
		if err != nil {
			http.Error(rw, err.Error(), 400)
		}
	}
	log.Print("DS", ds)
	jstring, err := json.Marshal(ds)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	rw.Write(jstring)
}

func DeleteDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice_mgo.DiceSet
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		log.Print(err)
		http.Error(rw, err.Error(), 400)
	}
	err = ds.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
	return
}
