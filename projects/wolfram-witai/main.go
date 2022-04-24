package main

import (
	// core packages
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// custom packages
	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load(".env")

	// slack bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// wit-ai client
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	// wolfram client
	wolframClient := &wolfram.Client{
		AppID: os.Getenv("WOLFRAM_APP_ID"),
	}

	go printCommandEvents(bot.CommandEvents())

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfram",
		Example:     "who is the president of Moldova",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")
			fmt.Println(query)

			response.Reply("Received")

			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			fmt.Println(msg)

			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])

			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			fmt.Println(value)

			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("There is an error")
			}
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
