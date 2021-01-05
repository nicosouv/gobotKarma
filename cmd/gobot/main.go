package main

import (
	Git "app/cmd/gobot/gitlab"
	Lib "app/cmd/gobot/lib"
	IM "app/cmd/gobot/slack"
	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/http"
	"sort"
	"strings"
)

var api = slack.New("SLACK_TOKEN")


type Success struct {
	Status int `json:"status"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/merge-requests", mergeRequestHandler).Methods("POST")
	router.HandleFunc("/merge-requests", commandHandler).Methods("GET")
	router.HandleFunc("/", commandHandler).Methods("POST", "GET")

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
	_ = r.ParseForm()
	user_opt_text := r.Form.Get("text")
	opts := strings.Fields(user_opt_text)

	var mrs []*gitlab.MergeRequest

	switch opts[0] {
		case "user":
			mrs = Git.GrabMRsForUsername(opts[1])
			break
		default:
			mrs = Git.GrabMRsForAllProjects("opened")
	}

	messages := []IM.SlackMessage{}

	sort.Slice(mrs, func(p, q int) bool {
		return mrs[p].ProjectID < mrs[q].ProjectID })

	for _, mr := range mrs {
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