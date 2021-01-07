package vcs

import (
	IM "app/cmd/gobot/im/slack"
	VCS "app/cmd/gobot/vcs/gitlab"
	"github.com/gorilla/mux"
	"github.com/xanzy/go-gitlab"
	"net/http"
	"sort"
	"strings"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET", "POST").Path("/merge-requests").HandlerFunc(mergeRequestHandler)
	router.Methods("GET", "POST").Path("/").HandlerFunc(commandHandler)

	return router
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
		mrs = VCS.GrabMRsForUsername(opts[1])
	default:
		mrs = VCS.GrabMRsForAllProjects(VCS.State_opened)
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