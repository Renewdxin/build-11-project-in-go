package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6240771686464-6220305923332-VHL7z9pfckBBcMDeVRu8mT0k")
	os.Setenv("CHANNEL_ID", "C066749D34N")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A065Z9DTPP1-6217853580642-d2e3b3e936101a355b8987adaaef7f3701234997cbdad1124484254a2ce3c39d")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"usingthymeleaf.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("name:%s, URL:%s\n", file.Name, file.URL)
	}
}
