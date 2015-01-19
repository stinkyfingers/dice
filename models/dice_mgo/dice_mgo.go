package dice_mgo

import (
	"github.com/stinkyfingers/dice/helpers/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"
	"math/rand"
)

type Die struct {
	Sides Sides `bson:"sides,omitempty" json:"sides, omitempty"`
}

type Dice []Die

type Side struct {
	Value string `bson:"value,omitempty" json:"value, omitempty"`
}
type Sides []Side

type DiceSet struct {
	ID bson.ObjectId `bson:"_id,omitempty"json:"id"`
	// ID       int           `bson:"id,omitempty" json:"id,omitempty"`
	Name   string        `bson:"name,omitempty" json:"name, omitempty"`
	Dice   Dice          `bson:"dice,omitempty" json:"dice, omitempty"`
	UserID bson.ObjectId `bson:"userId,omitempty" json:"userId, omitempty"`
	Public bool          `bson:"public,omitempty" json:"public, omitempty"`
}

type DiceSets []DiceSet

type Result struct {
	Value string `json:"value, omitempty"`
}
type Results []Result

func (ds *DiceSet) Roll() (Results, error) {
	var err error
	var r Result
	var rs Results
	err = ds.Get()
	if err != nil {
		return rs, err
	}
	for _, die := range ds.Dice {
		n := rand.Intn(len(die.Sides))
		r.Value = die.Sides[n].Value
		rs = append(rs, r)
	}
	return rs, err
}

func (ds *DiceSet) Create() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	ds.ID = bson.NewObjectId()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	err = c.Insert(ds)
	if err != nil {
		return err
	}
	return err
}

func (ds *DiceSet) Get() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	err = c.FindId(ds.ID).One(&ds)
	if err != nil {
		return err
	}
	return err
}

func GetUserDiceSets(userID bson.ObjectId) ([]DiceSet, error) {
	var err error
	var dss []DiceSet

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return dss, err
	}
	defer session.Close()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	err = c.Find(bson.M{"userId": userID}).All(&dss)
	if err != nil {
		return dss, err
	}
	return dss, err
}

func GetPublicDiceSets() ([]DiceSet, error) {
	var err error
	var dss []DiceSet
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		log.Print(err)
		return dss, err
	}
	defer session.Close()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	err = c.Find(bson.M{"public": true}).All(&dss)
	if err != nil {
		return dss, err
	}
	return dss, err
}

func (ds *DiceSet) Update() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	var ch mgo.Change
	ch.ReturnNew = true
	ch.Update = ds
	_, err = c.FindId(ds.ID).Apply(ch, &ds)
	if err != nil {
		return err
	}
	return err
}

func (ds *DiceSet) Delete() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB(database.MongoDatabase()).C("diceSets")
	err = c.Remove(bson.M{"_id": ds.ID})
	if err != nil {
		return err
	}
	return err
}
