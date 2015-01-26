package user_mgo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"github.com/stinkyfingers/dice/helpers/email"
	"github.com/stinkyfingers/dice/models/dice"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	// ID        int           `bson:"id,omitempty" json:"id,omitempty"`
	Email     string        `bson:"email,omitempty" json:"email,omitempty"`
	Password  string        `bson:"password,omitempty" json:"password,omitempty"`
	Encrypted string        `bson:"encrypted,omitempty" json:"encrypted,omitempty"`
	DiceSets  dice.DiceSets `bson:"diceSets,omitempty" json:"diceSets,omitempty"`
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

	u.ObjectID = bson.NewObjectId()
	c := session.DB(database.MongoDatabase()).C("users")
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

	c := session.DB(database.MongoDatabase()).C("users")
	err = c.FindId(u.ObjectID).One(&u)
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

	c := session.DB(database.MongoDatabase()).C("users")
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

	c := session.DB(database.MongoDatabase()).C("users")
	err = c.Remove(bson.M{"_id": u.ObjectID})
	if err != nil {
		return err
	}
	return err
}

func (u *User) ResetPassword() error {
	var err error
	var doc bytes.Buffer

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(database.MongoDatabase()).C("users")
	err = c.FindId(bson.M{"email": u.Email}).One(&u)
	if err != nil {
		return err
	}
	u.Password = randomPassword()
	if err != nil {
		return err
	}

	t := template.Must(template.New("email").Parse("Your new password is: " + u.Password + "."))
	t.Execute(&doc, nil)
	err = email.Send([]string{u.Email}, "Dice-A-Roni password reset.", doc)
	if err != nil {
		return err
	}
	u.Password = ""

	h := md5.New()
	h.Write([]byte(u.Password))
	u.Encrypted = hex.EncodeToString(h.Sum(nil))

	err = c.Update(bson.M{"_id": u.ObjectID}, bson.M{"$set": bson.M{"encrypted": u.Encrypted}})
	if err != nil {
		return err
	}

	return err
}

func randomPassword() string {
	var letters = []rune("qwertyuiopasdfghjklzxcvbnbm")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
