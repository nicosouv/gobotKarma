package slack

import (
	Lib "app/cmd/gobot/lib"
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SlackMessage struct {
	Title string
	Wip bool
	Branch string
	Upvote int
	Downvote int
	Author string
}

type Success struct {
	Status int `json:"status"`
}

func DisplayGitlabMRs(messages []SlackMessage) {
	var text strings.Builder
	for _, msg := range messages {
		text.WriteString("*" + msg.Branch + "* - " + msg.Title + " (_" + msg.Author + "_)" +
			"       :thumbsup: " + strconv.Itoa(msg.Upvote) + " / :thumbsdown: " + strconv.Itoa(msg.Downvote) + "\n")
	}

	data := url.Values{}
	data.Set("token", Lib.Getenv("SLACK_TOKEN"))
	data.Add("channel", Lib.Getenv("SLACK_CHANNEL"))
	data.Add("text", text.String())

	text.Reset()

	log.Println("sending Message to Slack")
	_, err := http.Post(
		"https://slack.com/api/chat.postMessage",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}
}


func NoIdea() {
	data := url.Values{}
	data.Set("token", Lib.Getenv("SLACK_TOKEN"))
	data.Add("channel", Lib.Getenv("SLACK_CHANNEL"))
	data.Add("text", "I have no idea what you want.")

	log.Println("sending Message to Slack")
	_, err := http.Post(
		"https://slack.com/api/chat.postMessage",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}
}
