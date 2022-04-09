package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp )
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("")
	}
}

func main(){
	fmt.Println("Hello")

	os.Setenv("SLACK_BOT_TOKEN","xoxb-333529780928-3331611353527-24EXkBnyW1dcabZZS9g3mztX")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A03AVSVM42C-3331566023591-bcf8f4a166230d9f0a4c4c9e34c55cf540c6f415e22a3cdb01b2caaa5f236c39")

	bot:=slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year:=request.Param("year")

			yob, err := strconv.Atoi(year)
			if err!=nil{
				println("error")
			}
			age:=2021-yob

			r:=fmt.Sprintf("Age is %d", age)
			fmt.Println(r)
			response.Reply(r)
		},
	})


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	err := bot.Listen(ctx)
	if err!=nil{
		log.Fatal(err)
	}
}
