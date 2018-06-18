package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Slack slack integration
type Slack struct {
	url string
}

// Message request
type Message struct {
	Attachments []Attache `json:"attachments,omitempty"`
}

// Attache request attache
type Attache struct {
	Text      string `json:"text,omitempty"`
	PreText   string `json:"pretext,omitempty"`
	Author    string `json:"author_name,omitempty"`
	Title     string `json:"title,omitempty"`
	Footer    string `json:"footer,omitempty"`
	Color     string `json:"color,omitempty"`
	TitleLink string `json:"title_link,omitempty"`
	Actions   interface{}
}

//SendMessage  send message to slack
func (s Slack) SendMessage(message Message) {
	jsonBytes, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	log.Printf("Sending body: %v", string(jsonBytes))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Println("Slack Response: " + bodyString)
}
