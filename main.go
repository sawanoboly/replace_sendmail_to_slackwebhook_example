package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
)

var webhookURL = "https://hooks.slack.com/services/foo/bar/baz"

func main() {
	sendmainInput, _ := ioutil.ReadAll(os.Stdin)
	// fmt.Println(string(sendmainInput))
	m, err := mail.ReadMessage(strings.NewReader(string(sendmainInput)))

	if err != nil {
		log.Fatal(err)
	}
	header := m.Header
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: header.Get("Subject"), Value: string(body)})

	username, _ := os.Hostname()
	payload := slack.Payload{
		Username:    username,
		Text:        string(body),
		Attachments: []slack.Attachment{attachment1},
	}

	errs := slack.Send(webhookURL, "", payload)
	if len(errs) > 0 {
		fmt.Printf("error: %s\n", errs)
	}
}
