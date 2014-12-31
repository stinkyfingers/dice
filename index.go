package main

import (
	"flag"
	"fmt"
	"github.com/stinkyfingers/dice/controllers/application"
	// "github.com/stinkyfingers/dice/controllers/dice"
	// "github.com/stinkyfingers/dice/controllers/user"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	port = flag.String("port", ":8080", "Port to run on")
)

func main() {
	flag.Parse()
	fmt.Print("Dice Running. \n")

	// http.Handle("/public/js/", http.StripPrefix("/public/js/", http.FileServer(http.Dir("public/js"))))
	// http.Handle("/public/templates/", http.StripPrefix("/public/templates/", http.FileServer(http.Dir("public/templates"))))

	// http.HandleFunc("/roll", dice.Roll)
	// http.HandleFunc("/getDiceSet", dice.GetDiceSet)
	// http.HandleFunc("/getDie", dice.GetDie)
	// http.HandleFunc("/getSide", dice.GetSide)
	// http.HandleFunc("/saveDiceSet", dice.SaveDiceSet)
	// http.HandleFunc("/saveDie", dice.SaveDie)
	// http.HandleFunc("/saveSide", dice.SaveSide)
	// http.HandleFunc("/deleteDiceSet", dice.DeleteDiceSet)
	// http.HandleFunc("/deleteDie", dice.DeleteDie)
	// http.HandleFunc("/deleteSide", dice.DeleteSide)
	// http.HandleFunc("/app", application.Application)
	// http.HandleFunc("/test", application.Application)
	//TODO Roll your own regex handler - http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc

	rh.AddRoute(regexp.MustCompile("/public/js/"), http.StripPrefix("/public/js/", http.FileServer(http.Dir("public/js"))))
	rh.AddRoute(regexp.MustCompile("/public/templates/"), http.StripPrefix("/public/templates/", http.FileServer(http.Dir("public/templates"))))

	rh.AddRoute(regexp.MustCompile("/test"), http.HandlerFunc(application.Application))
	rh.AddRoute(regexp.MustCompile("/app"), http.HandlerFunc(application.Application))
	rh.AddRoute(regexp.MustCompile("/.*"), http.HandlerFunc(application.Application))
	log.Print(rh.routes)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), &rh)
	if err != nil {
		log.Print(err)
	}
}

var rh RegexpHandler

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}
type RegexpHandler struct {
	routes []*route
}

// func AddRoute(pattern *regexp.Regexp, handler http.Handler) *RegexpHandler {
// 	var rh RegexpHandler
// 	ro := route{pattern: pattern, handler: handler}
// 	rh.routes = append(rh.routes, &ro)
// 	return &rh
// }

func (rh *RegexpHandler) AddRoute(pattern *regexp.Regexp, handler http.Handler) {
	ro := route{pattern: pattern, handler: handler}
	rh.routes = append(rh.routes, &ro)
}

func (rh *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rh.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
