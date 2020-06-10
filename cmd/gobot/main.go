package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"encoding/json"

	Lib "./lib"

	slack "github.com/slack-go/slack"
)

type Success struct {
	Status int `json:"status"`
}

func main() {
	http.HandleFunc("/", commandHandler)

	port := Lib.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}


func commandHandler(w http.ResponseWriter, r *http.Request) {
	data := url.Values{}
	data.Set("token", Lib.Getenv("SLACK_TOKEN"))
	data.Add("channel", Lib.Getenv("SLACK_CHANNEL"))
	data.Add("text", "yo")

	log.Println("sending Message to Slack")
	resp, err := http.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	} else {
		success := Success{resp.StatusCode}
		jsonized, e := json.Marshal(success)

		if e != nil {
			log.Fatal(e)
		}

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(jsonized))
	}


	// testing stuff
	api := slack.New(Lib.Getenv("SLACK_TOKEN"))
	user, err := api.GetUserInfo("U043WHXCV")
	if err != nil {
		log.Println("%s\n", err)
	} else {
		log.Println("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
	}
}