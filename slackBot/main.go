package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// Function to print command events
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println("Timestamp:", event.Timestamp)
		fmt.Println("Command:", event.Command)
		fmt.Println("Parameters:", event.Parameters)
		fmt.Println("Event:", event.Event)
		fmt.Println()
	}
}

func main() {
	// Set environment variables (move sensitive values to a .env file for security)
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	// Initialize Slack bot client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Start listening for command events asynchronously
	go printCommandEvents(bot.CommandEvents())

	// Define bot command
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year format. Please provide a valid number.")
				return
			}
			age := 2025 - yob
			response.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})

	// Create context to manage bot lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start bot listener
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
