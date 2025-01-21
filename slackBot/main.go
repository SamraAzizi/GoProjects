package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func main() {

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOEKN"))
}
fmt"aspdgh
