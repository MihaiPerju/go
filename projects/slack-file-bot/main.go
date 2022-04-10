package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-333529780928-3393191378688-iSaY6jiPYh7o1ewUFTWP70fi")
	os.Setenv("SLACK_CHANNEL_ID", "C9UB51Q2Z")

	api:=slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}

	fileArr:=[]string{"test-file.txt"}

	for i := 0;i<len(fileArr);i++{
		params:=slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}


		file, err:=api.UploadFile(params)
		if err!=nil{
			fmt.Printf("%s\n", err)
			return 
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}