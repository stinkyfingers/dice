package database

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
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

// func MongoConnectionString() *mgo.DialInfo {
// 	var info mgo.DialInfo
// 	addr := os.Getenv("MONGO_URL")
// 	if addr == "" {
// 		addr = "127.0.0.1"
// 	}

// 	info.Addrs = append(info.Addrs, addr)
// 	info.Username = os.Getenv("MONGO_CART_USERNAME")
// 	info.Password = os.Getenv("MONGO_CART_PASSWORD")
// 	info.Database = os.Getenv("MONGO_CART_DATABASE")
// 	info.Timeout = time.Second * 2
// 	if info.Database == "" {
// 		info.Database = "Mullets"
// 	}

// 	return &info
// }
