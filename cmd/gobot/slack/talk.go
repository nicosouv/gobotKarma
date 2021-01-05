package slack

import (
	Lib "app/cmd/gobot/lib"
	"bytes"
	"log"
	"math/rand"
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


func sendMessageToSlack(text string) {
	if text == "" {
		log.Printf("[WARNING] We don't send empty text to Slack")
		return
	}

	data := url.Values{}
	data.Set("token", Lib.Getenv("SLACK_TOKEN"))
	data.Add("channel", Lib.Getenv("SLACK_CHANNEL"))
	data.Add("text", text)

	log.Println("sending message to Slack")
	_, err := http.Post(
		"https://slack.com/api/chat.postMessage",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func DisplayGitlabMRs(messages []SlackMessage) {
	var text strings.Builder

	text.WriteString("┌────────────────┐" + "\n" +
		                " │          *Merge Requests*         │" + "\n" +
		                "└────────────────┘" + "\n\n")

	text.WriteString("Hey,\nVoici les merge requests en attente!\n\n")

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

	sendMessageToSlack(text.String())
}


func NoIdea() {
	message := [5]string{
		"I have no idea what you want.",
		"Hey, leave me alone",
		"Stop bothering me",
		"Oh god, what do you want again?!",
		"k, thx, bye",
	}

	n := rand.Intn(len(message)-1)

	sendMessageToSlack(message[n])
}