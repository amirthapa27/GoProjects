package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "Enter-your-slack-bot-token")
	os.Setenv("CHANNEL_ID", "Enter your channel ID")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN")) //creating  new connection
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"sample.pdf"} //name of the file to be uploaded

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(file.Name, file.URL)
	}
}
