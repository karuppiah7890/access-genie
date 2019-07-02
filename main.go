package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func main() {
	slackToken, envExists := os.LookupEnv("SLACK_TOKEN")
	if !envExists {
		panic("error: SLACK_TOKEN environment variable is not set")
	}

	bot := slacker.NewClient(slackToken)

	pingCommand := &slacker.CommandDefinition{
		Description: "Ping!",
		Example:     "ping",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", pingCommand)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
