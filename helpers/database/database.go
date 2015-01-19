package database

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	// "log"
	"os"
	"time"
)

var (
	EmptyDb = flag.String("clean", "", "bind empty database with structure defined")
)

func ConnectionString() string {

	if addr := os.Getenv("CLEARDB_DATABASE_URL"); addr != "" {
		return fmt.Sprint(addr)
	}

	if EmptyDb != nil && *EmptyDb != "" {
		return "root:@tcp(127.0.0.1:3306)/Dice_Empty?parseTime=true&loc=America%2FChicago"
	}
	return "root:@tcp(127.0.0.1:3306)/Dice?parseTime=true&loc=America%2FChicago"
}

func MongoConnectionString() *mgo.DialInfo {

	var (
		MongoDBHosts    = os.Getenv("DBHOST")
		AuthDatabase    = os.Getenv("MONGO_DB")
		AuthUserName    = os.Getenv("MONGO_USER")
		AuthPassword    = os.Getenv("MONGO_PASS")
		mongoDBDialInfo mgo.DialInfo
	)

	if MongoDBHosts == "" {
		mongoDBDialInfo = mgo.DialInfo{
			Addrs: []string{"127.0.0.1"},
		}
	} else {
		mongoDBDialInfo = mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Timeout:  60 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}
	}

	return &mongoDBDialInfo
}

func MongoDatabase() string {
	var db string
	db = os.Getenv("MONGO_DB")
	if db == "" {
		db = "wilddice"
	}

	return db
}
