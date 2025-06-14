package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := string(os.Getenv("CHANNEL_ID"))
	fileArr := []string{"test.txt"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileV2Parameters{
			Channel:  channelArr,
			File:     fileArr[i],
			Filename: fileArr[i],
			FileSize: 40,
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s", file.Title)
	}
}
