package slack

import (
	Lib "app/cmd/gobot/lib"
	"bytes"
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"net/http"
	"net/url"
)


type Success struct {
	Status int `json:"status"`
}

func DisplayGitlabMRs(data string) {
	api := slack.New("SLACK_TOKEN")
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage(
		"SLACK_CHANNEL_ID",
		slack.MsgOptionText(data, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
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
