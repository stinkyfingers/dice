package dice_mgo

import (
	"github.com/stinkyfingers/dice/helpers/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"math/rand"
)

type Die struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	DiceSetID bson.ObjectId `bson:"diceSetId,omitempty"`
	Sides     Sides         `bson:"sides,omitempty"`
}

type Dice []Die

type Side struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	DieID bson.ObjectId `bson:"dieId,omitempty"`
	Value string        `bson:"value,omitempty"`
}
type Sides []Side

type DiceSet struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `bson:"name,omitempty"`
	Dice   Dice          `bson:"dice,omitempty"`
	UserID bson.ObjectId `bson:"userId,omitempty"`
	Public bool          `bson:"public,omitempty"`
}

type DiceSets []DiceSet

func (d *Die) Roll() (string, error) {
	var err error
	err = d.Get()
	if err != nil {
		return "", err
	}
	n := rand.Intn(len(d.Sides))
	return d.Sides[n].Value, err
}

func (s *Side) Create() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	s.ID = bson.NewObjectId()
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
	d.ID = bson.NewObjectId()
	c := session.DB("wilddice").C("dice")
	err = c.Insert(d)
	if err != nil {
		return err
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
	ds.ID = bson.NewObjectId()
	c := session.DB("wilddice").C("diceSets")
	err = c.Insert(ds)
	if err != nil {
		return err
	}
	return err
}

func (s *Side) Get() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	err = c.FindId(s.ID).One(&s)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Get() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	err = c.FindId(d.ID).One(&d)
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
	c := session.DB("wilddice").C("diceSets")
	err = c.FindId(ds.ID).One(&ds)
	if err != nil {
		return err
	}
	return err
}

func (s *Side) Delete() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	err = c.Remove(bson.M{"_id": s.ID})
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Delete() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	err = c.Remove(bson.M{"_id": d.ID})
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
	c := session.DB("wilddice").C("diceSets")
	err = c.Remove(bson.M{"_id": ds.ID})
	if err != nil {
		return err
	}
	return err
}
