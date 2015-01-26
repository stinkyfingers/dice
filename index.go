package main

import (
	"flag"
	"fmt"
	"github.com/stinkyfingers/dice/controllers/application"
	"github.com/stinkyfingers/dice/controllers/dice"
	"github.com/stinkyfingers/dice/controllers/user"
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

	//FILES
	rh.AddRoute(regexp.MustCompile("/public/js/"), http.StripPrefix("/public/js/", http.FileServer(http.Dir("public/js"))))
	rh.AddRoute(regexp.MustCompile("/public/templates/"), http.StripPrefix("/public/templates/", http.FileServer(http.Dir("public/templates"))))
	rh.AddRoute(regexp.MustCompile("/public/css/"), http.StripPrefix("/public/css/", http.FileServer(http.Dir("public/css"))))

	//API
	rh.AddRoute(regexp.MustCompile("/roll"), http.HandlerFunc(dice.Roll))
	rh.AddRoute(regexp.MustCompile("/getPublicDiceSets"), http.HandlerFunc(dice.GetPublicDiceSets))
	rh.AddRoute(regexp.MustCompile("/resetPassword"), http.HandlerFunc(user.ResetPassword))
	rh.AddRoute(regexp.MustCompile("/getUserDiceSets"), http.HandlerFunc(dice.GetUserDiceSets))
	rh.AddRoute(regexp.MustCompile("/getDiceSet"), http.HandlerFunc(dice.GetDiceSet))
	rh.AddRoute(regexp.MustCompile("/saveDiceSet"), http.HandlerFunc(dice.SaveDiceSet))
	rh.AddRoute(regexp.MustCompile("/deleteDiceSet"), http.HandlerFunc(dice.DeleteDiceSet))

	//ROUTES
	rh.AddRoute(regexp.MustCompile("/login"), http.HandlerFunc(application.Login))
	rh.AddRoute(regexp.MustCompile("/logout"), http.HandlerFunc(application.Logout))
	rh.AddRoute(regexp.MustCompile("/auth"), http.HandlerFunc(user.AuthenticateUser))
	rh.AddRoute(regexp.MustCompile("/test"), http.HandlerFunc(application.Application))
	rh.AddRoute(regexp.MustCompile("/app"), http.HandlerFunc(application.Application))
	rh.AddRoute(regexp.MustCompile("/register"), http.HandlerFunc(user.Register))
	rh.AddRoute(regexp.MustCompile("/.*"), http.HandlerFunc(application.Application))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), &rh)
	if err != nil {
		log.Print(err)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request) string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fn(rw, r)
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
