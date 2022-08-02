package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type mems struct {
	id string
}


func GetMembers() (map[string]interface{}, error) {

    token := os.Getenv("SLACK_AUTH_TOKEN")
    bearer := "Bearer " + token
    url := "https://slack.com/api/users.list"
    
    
    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)
    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
            log.Println("Error while reading the response bytes:", err)
    }

	

    m := make(map[string]interface{})
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
 	   log.Fatal(err)
	}
	
	var members  = m["members"]
	fmt.Println("m!!!", m)
	fmt.Println("members!!!", members)
	return m, err1; 
}

func main() {
    godotenv.Load(".env")

    token := os.Getenv("SLACK_AUTH_TOKEN")
    channelID := os.Getenv("SLACK_CHANNEL_ID")


    client := slack.New(token, slack.OptionDebug(true))
    // client.GetBotInfo()
    
    abc, _ := GetMembers()

    fmt.Println("ress!!", abc)

	// Create the Slack attachment that we will send to the channel
	attachment := slack.Attachment{
		Pretext: "Super Bot Message",
		Text:    "some text",
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: "#36a64f",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

    

	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, _, err := client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		// slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)
if err != nil {
		panic(err)
	}
	// fmt.Printf("Message sent at %s", timestamp)
    
}