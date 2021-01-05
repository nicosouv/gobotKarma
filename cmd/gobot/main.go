package main

import (
	Git "app/cmd/gobot/gitlab"
	Lib "app/cmd/gobot/lib"
	Slack "app/cmd/gobot/slack"
	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
	"log"
	"net/http"
)

var api = slack.New("SLACK_TOKEN")


type Success struct {
	Status int `json:"status"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/merge-requests", mergeRequestHandler).Methods("POST")
	router.HandleFunc("/merge-requests", mergeRequestHandler).Methods("GET")
	router.HandleFunc("/", commandHandler).Methods("POST")
	router.HandleFunc("/", commandHandler).Methods("GET")

	port := Lib.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	Slack.NoIdea()
}

func mergeRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("MR Handler")

	Git.GrabMRsForAllProjects()

	log.Printf("Done")
	/*for _, mr := range mrs {
		fmt.Print("%v\n", mr)
	}*/

	//Slack.DisplayGitlabMRs("HEY")
}