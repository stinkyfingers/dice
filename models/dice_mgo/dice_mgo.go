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

func (ds *DiceSet) GetDiceByDiceSetID() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	err = c.FindId(bson.M{"diceSetId": ds.ID}).All(&ds.Dice)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) GetSidesByDiceID() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	err = c.FindId(bson.M{"diceId": d.ID}).All(&d.Sides)
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
	c := session.DB("wilddice").C("diceSets")
	err = c.FindId(bson.M{"userId": userID}).All(&dss)
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
		return dss, err
	}
	defer session.Close()
	c := session.DB("wilddice").C("diceSets")
	err = c.FindId(bson.M{"isPublic": true}).All(&dss)
	if err != nil {
		return dss, err
	}
	return dss, err
}

func (s *Side) Update() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	var ch mgo.Change
	ch.ReturnNew = true
	ch.Update = s
	_, err = c.Find(bson.M{"_id": s.ID}).Apply(ch, &s)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Update() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	var ch mgo.Change
	ch.ReturnNew = true
	ch.Update = d
	_, err = c.Find(bson.M{"_id": d.ID}).Apply(ch, &d)
	if err != nil {
		return err
	}
	return err
}

func (ds *DiceSet) Update() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("diceSets")
	var ch mgo.Change
	ch.ReturnNew = true
	ch.Update = ds
	_, err = c.Find(bson.M{"_id": ds.ID}).Apply(ch, &ds)
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

func (d *Die) DeleteSides() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("sides")
	err = c.Remove(bson.M{"dieId": d.ID})
	if err != nil {
		return err
	}
	return err
}

func (ds *DiceSet) DeleteDice() error {
	var err error
	for _, d := range ds.Dice {
		err = d.DeleteSides()
		if err != nil {
			return err
		}
	}
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("wilddice").C("dice")
	err = c.Remove(bson.M{"diceSet": ds.ID})
	if err != nil {
		return err
	}
	return err
}
