package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/karuppiah7890/access-genie/slackutils"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
)

func main() {
	slackToken, envExists := os.LookupEnv("SLACK_TOKEN")
	if !envExists {
		panic("error: SLACK_TOKEN environment variable is not set")
	}

	debugStr := os.Getenv("DEBUG")
	debug, _ := strconv.ParseBool(debugStr)

	bot := slacker.NewClient(slackToken, slacker.WithDebug(debug))

	pingCommand := &slacker.CommandDefinition{
		Description: "Ping!",
		Example:     "ping",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	listCommand := &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			slackMessage := slackutils.SelectBlockForList([]string{"cluster-one", "cluster-two"})
			section := slackutils.SectionBlock("Choose the cluster you want access to", slackMessage)
			blocks := []slack.Block{section}
			response.Reply("", slacker.WithBlocks(blocks))
		},
	}

	bot.Command("ping", pingCommand)
	bot.Command("list", listCommand)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
