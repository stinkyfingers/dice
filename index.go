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

	http.HandleFunc("/roll", dice.Roll)
	http.HandleFunc("/getDiceSet", dice.GetDiceSet)
	http.HandleFunc("/getDie", dice.GetDie)
	http.HandleFunc("/getSide", dice.GetSide)
	http.HandleFunc("/saveDiceSet", dice.SaveDiceSet)
	http.HandleFunc("/saveDie", dice.SaveDie)
	http.HandleFunc("/saveSide", dice.SaveSide)
	http.HandleFunc("/deleteDiceSet", dice.DeleteDiceSet)
	http.HandleFunc("/deleteDie", dice.DeleteDie)
	http.HandleFunc("/deleteSide", dice.DeleteSide)

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
