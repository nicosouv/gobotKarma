package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"encoding/json"
)

func main() {

	http.HandleFunc("/", commandHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type Success struct {
	Status int `json:"status"`
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}


func commandHandler(w http.ResponseWriter, r *http.Request) {
	data := url.Values{}
	data.Set("token", getenv("SLACK_TOKEN"))
	data.Add("channel", getenv("SLACK_CHANNEL"))
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
}