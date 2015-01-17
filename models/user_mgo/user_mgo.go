package user_mgo

import (
	"crypto/md5"
	"encoding/hex"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"github.com/stinkyfingers/dice/models/dice"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	Encrypted string        `bson:"encrypted"`
	DiceSets  dice.DiceSets `bson:"diceSets"`
}

func (u *User) Create() error {
	var err error
	h := md5.New()
	h.Write([]byte(u.Password))
	u.Encrypted = hex.EncodeToString(h.Sum(nil))
	u.Password = ""

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	u.ID = bson.NewObjectId()
	c := session.DB("wilddice").C("users")
	err = c.Insert(u)
	if err != nil {
		return err
	}
	return err
}

func (u *User) Get() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("wilddice").C("users")
	err = c.FindId(u.ID).One(&u)
	if err != nil {
		return err
	}

	return err
}
func (u *User) Authenticate() error {
	var err error
	h := md5.New()
	h.Write([]byte(u.Password))
	u.Encrypted = hex.EncodeToString(h.Sum(nil))
	u.Password = ""

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("wilddice").C("users")
	err = c.Find(bson.M{"email": u.Email, "encrypted": u.Encrypted}).One(&u)
	if err != nil {
		return err
	}

	return err
}

func (u *User) Delete() error {
	var err error
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("wilddice").C("users")
	err = c.Remove(bson.M{"_id": u.ID})
	if err != nil {
		return err
	}
	return err
}
