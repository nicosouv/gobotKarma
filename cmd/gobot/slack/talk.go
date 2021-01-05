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

	text.WriteString("┌────────────────┐" + "\n" +
		                " │          *Merge Requests*         │" + "\n" +
		                "└────────────────┘" + "\n\n")

	for _, msg := range messages {
		str := "*" + msg.Branch + "* - " + msg.Title + " (_" + msg.Author + "_)" + "\n"
		str = str + "\n>  `⭡` "

		if msg.Upvote > 0 {
			str += "*" + strconv.Itoa(msg.Upvote) + "*"
		} else {
			str += strconv.Itoa(msg.Upvote)
		}

		str += " / `⭣` "

		if msg.Downvote > 0 {
			str += "*" + strconv.Itoa(msg.Downvote) + "*"
		} else {
			str += strconv.Itoa(msg.Downvote)
		}

		str += "\n"

		text.WriteString(str)
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
