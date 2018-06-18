package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var slackURL string
	var upsourceURL string
	flag.StringVar(&slackURL, "slack", os.Getenv("SLACK_URL"), "a string")
	flag.StringVar(&upsourceURL, "upsource", os.Getenv("UPSOURCE_URL"), "a string")
	flag.Parse()
	slack := Slack{url: slackURL}
	http.HandleFunc("/slack", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		log.Printf("Reading body: %v", string(body))
		result := Parse(body, upsourceURL)
		message := AdaptRequestResult(result)
		slack.SendMessage(message)
	})
	if err := http.ListenAndServe(":8989", nil); err != nil {
		panic(err)
	}
}
