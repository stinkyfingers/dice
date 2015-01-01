package dice

import (
	"encoding/json"
	"github.com/stinkyfingers/dice/models/dice"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type result struct {
	DieID int
	Value string
}

type results []result

func Roll(rw http.ResponseWriter, r *http.Request) {
	var ds dice.DiceSet
	var rs results
	var thisDie result

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = ds.Get()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	for _, d := range ds.Dice {
		r, err := d.Roll()
		if err != nil {
			http.Error(rw, err.Error(), 404)
		}
		thisDie.DieID = d.ID
		thisDie.Value = r
		rs = append(rs, thisDie)
	}

	jstring, err := json.Marshal(rs)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func GetPublicDiceSets(rw http.ResponseWriter, r *http.Request) {
	var dss []dice.DiceSet
	var err error

	dss, err = dice.GetPublicDiceSets()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	jstring, err := json.Marshal(dss)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func GetUserDiceSets(rw http.ResponseWriter, r *http.Request) {
	var dss []dice.DiceSet
	var err error

	cookie, err := r.Cookie("user")
	if err != nil || cookie == nil {
		http.Error(rw, err.Error(), 404)
	}
	c := strings.Split(cookie.String(), "=")[1]
	userId, err := strconv.Atoi(c)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	dss, err = dice.GetUserDiceSets(userId)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	jstring, err := json.Marshal(dss)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jstring)
}

func GetDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice.DiceSet

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = ds.Get()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	jstring, err := json.Marshal(ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func GetDie(rw http.ResponseWriter, r *http.Request) {
	var d dice.Die

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &d)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = d.Get()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	jstring, err := json.Marshal(d)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}
func GetSide(rw http.ResponseWriter, r *http.Request) {
	var s dice.Side

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &s)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = s.Get()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	jstring, err := json.Marshal(s)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}
func SaveDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice.DiceSet

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	if ds.ID > 0 {
		err = ds.Get()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
		err = ds.Update()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	} else {
		err = ds.Create()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	}
	jstring, err := json.Marshal(ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func SaveDie(rw http.ResponseWriter, r *http.Request) {
	var d dice.Die

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &d)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	if d.ID > 0 {
		err = d.Get()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
		err = d.Update()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	} else {
		err = d.Create()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	}
	jstring, err := json.Marshal(d)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func SaveSide(rw http.ResponseWriter, r *http.Request) {
	var s dice.Side

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &s)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	if s.ID > 0 {
		err = s.Get()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
		err = s.Update()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	} else {
		err = s.Create()
		if err != nil {
			http.Error(rw, err.Error(), 401)
		}
	}
	jstring, err := json.Marshal(s)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	rw.Write(jstring)
}

func DeleteDiceSet(rw http.ResponseWriter, r *http.Request) {
	var ds dice.DiceSet
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &ds)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = ds.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	return
}

func DeleteDie(rw http.ResponseWriter, r *http.Request) {
	var d dice.Die
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &d)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = d.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	return
}
func DeleteSide(rw http.ResponseWriter, r *http.Request) {
	var s dice.Side
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}

	err = json.Unmarshal(requestBody, &s)
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	err = s.Delete()
	if err != nil {
		http.Error(rw, err.Error(), 404)
	}
	return
}
