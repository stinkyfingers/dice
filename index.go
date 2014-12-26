package main

import (
	"flag"
	"fmt"
	"github.com/stinkyfingers/dice/controllers/dice"
	"log"
	"net/http"
	"os"
)

var (
	port = flag.String("port", ":8080", "Port to run on")
)

func main() {
	flag.Parse()
	fmt.Print("Dice Running. \n")

	// http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css"))))

	http.HandleFunc("/getDie", dice.GetDie)

	// if port != nil && *port != "" {
	// 	fmt.Print("Port: " + *port)
	// 	http.ListenAndServe(":"+*port, nil)
	// }
	// http.ListenAndServe(":8080", http.FileServer(http.Dir("/mullets/images"))) //Not needed, but catches port
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Print(err)
	}
}
