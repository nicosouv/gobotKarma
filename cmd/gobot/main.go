package main

import (
	Lib "app/cmd/gobot/lib"
	IM_Router "app/cmd/gobot/im"
	VCS_Router "app/cmd/gobot/vcs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)


type Success struct {
	Status int `json:"status"`
}

const event_route = "event"

func main() {
	router := mux.NewRouter()

	mount(router, "/im", IM_Router.InitializeRouter())
	mount(router, "", VCS_Router.InitializeRouter())

	port := Lib.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

// mount mux router to be handle in their specific directory
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}

