package main

import (
	Git "app/cmd/gobot/gitlab"
	Lib "app/cmd/gobot/lib"
	IM "app/cmd/gobot/slack"
	"fmt"
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
	IM.NoIdea()
}

func mergeRequestHandler(w http.ResponseWriter, r *http.Request) {

	var mrs = Git.GrabMRsForAllProjects()
	messages := []IM.SlackMessage{}

	for _, mr := range mrs {
		fmt.Print("%v\n", mr)
		msg := IM.SlackMessage{
			Title:    mr.Title,
			Wip:      mr.WorkInProgress,
			Branch:   mr.SourceBranch,
			Upvote:   mr.Upvotes,
			Downvote: mr.Downvotes,
			Author:   mr.Author.Name,
		}

		messages = append(messages, msg)
	}

	IM.DisplayGitlabMRs(messages)
}