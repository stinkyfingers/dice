package dice_mgo

import (
	"github.com/stinkyfingers/dice/helpers/database"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "math/rand"
)

type Die struct {
	ID        int   `json:"id,omitempty"`
	DiceSetID int   `json:"diceSetId,omitempty"`
	Sides     Sides `json:"sides,omitempty"`
}

type Dice []Die

type Side struct {
	ID    int    `json:"id,omitempty"`
	DieID int    `json:"dieId,omitempty"`
	Value string `json:"value,omitempty"`
}
type Sides []Side

type DiceSet struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Dice   Dice   `json:"dice,omitempty"`
	UserID int    `json:"userId,omitempty"`
	Public bool   `json:"public,omitempty"`
}

type DiceSets []DiceSet

func (s *Side) Create() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Create() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	err = c.Insert(d)
	if err != nil {
		return err
	}
	for _, s := range d.Sides {
		err = s.Create()
		if err != nil {
			return err
		}
	}
	return err
}

func (ds *DiceSet) Create() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("diceSet")
	err = c.Insert(ds)
	if err != nil {
		return err
	}
	for _, d := range ds.Dice {
		err = d.Create()
		if err != nil {
			return err
		}
		for _, s := range d.Sides {
			err = s.Create()
			if err != nil {
				return err
			}
		}
	}
	return err
}
