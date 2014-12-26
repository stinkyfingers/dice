package dice

import (
	"encoding/json"
	"github.com/stinkyfingers/dice/models/dice"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	DieID int
	Value string
}

type results []result

func GetDie(rw http.ResponseWriter, r *http.Request) {
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
