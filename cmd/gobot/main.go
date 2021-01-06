package main

import (
	Git "app/cmd/gobot/gitlab"
	Lib "app/cmd/gobot/lib"
	IM "app/cmd/gobot/slack"
	"github.com/gorilla/mux"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/http"
	"sort"
	"strings"
)


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

	if len(opts) == 0 {
		opts = append(opts, "")
	}

	switch opts[0] {
		case "wtf", "fuck", "test", "hello", "hey", "salut":
			commandHandler(w, r)
		case "user":
			mrs = Git.GrabMRsForUsername(opts[1])
		default:
			mrs = Git.GrabMRsForAllProjects(Git.State_opened)
	}

	if len(mrs) > 0 {
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
}