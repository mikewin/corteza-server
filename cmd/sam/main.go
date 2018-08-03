package main

import (
	"log"
	"net"
	"os"

	"net/http"

	"github.com/go-chi/chi"

	"github.com/crusttech/crust/auth"
	"github.com/crusttech/crust/rbac"
	"github.com/crusttech/crust/sam/rest"
	"github.com/crusttech/crust/sam/websocket"
	"github.com/titpetric/factory"
)

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func main() {
	config := flags("sam", auth.Flags, rbac.Flags, websocket.Flags)

	// log to stdout not stderr
	log.SetOutput(os.Stdout)
	go NewMonitor(config.monitorInterval)

	// set up database connection
	factory.Database.Add("default", config.dbDSN)
	db, err := factory.Database.Get()
	handleError(err, "Can't connect to database")
	db.Profiler = &factory.Database.ProfilerStdout

	// listen socket for http server
	log.Println("Starting http server on address " + config.httpAddr)
	listener, err := net.Listen("tcp", config.httpAddr)
	handleError(err, "Can't listen on addr "+config.httpAddr)

	// route options
	routeOptions, err := RouteOptions{}.New()
	handleError(err, "Error creating RouteOptions object")

	r := chi.NewRouter()

	// JWT Auth
	jwtAuth, err := auth.JWT()
	handleError(err, "Error creating JWT Auth object")
	r.Use(jwtAuth.Verifier(), jwtAuth.Authenticator())

	// mount REST & WS routes
	MountRoutes(r, routeOptions, rest.MountRoutes(jwtAuth), websocket.MountRoutes())

	http.Serve(listener, r)
}
