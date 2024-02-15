package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Usage() {
	fmt.Fprintln(os.Stdout, `Usage:
	tgbot-msg [Options] <msg>

Options:
	-h, --help		Show this output
	-chat-id=<value>	<value> specify where the message send to

Environment Variables:
	TGBOT_MSG_TOKEN		api token of telegram bot (required)
	TGBOT_MSG_CHATID	see -chat-id, at least one of them must provide
				-chat-id override TGBOT_MSG_CHATID when both provided`)
}

func fail(s string) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "Show the usage (short)")
	flag.BoolVar(&help, "help", false, "Show the usage")

	flagChatID := flag.String("chat-id", "-1", "chat ID specify where the message send to")
	flag.Parse()

	if help {
		Usage()
		os.Exit(0)
	}

	token, exists := os.LookupEnv("TGBOT_MSG_TOKEN")
	if !exists {
		Usage()
		fail("Error: env 'TGBOT_MSG_TOKEN' env is required")
	}

	var chatID int64
	var chatIDStr string
	chatIDStr, exists = os.LookupEnv("TGBOT_MSG_CHATID")
	if exists {
		num, err := strconv.ParseInt(chatIDStr, 10, 64)
		if err != nil {
			Usage()
			fail("Error: env 'TGBOT_MSG_CHATID' is not a valid integer")
		} else {
			chatID = num
		}
	}

	if chatID == 0 && *flagChatID == "-1" {
		Usage()
		fail("Error: Provide Chat ID either from env 'TGBOT_MSG_CHATID' or -chat-id")
	}
	if *flagChatID != "-1" {
		num, err := strconv.ParseInt(*flagChatID, 10, 64)
		if err != nil {
			Usage()
			fail("Error: flag -chat-id <value> is not a valid integer")
		} else {
			chatID = num
		}
	}

	args := flag.Args()
	if len(args) != 1 {
		Usage()
		fail(fmt.Sprint("Error: provide exact one positional argument for", os.Args[0]))
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fail(fmt.Sprint("Error: tgbotapi.NewBotAPI() ", err))
	}

	msg := tgbotapi.NewMessage(chatID, args[0])
	bot.Send(msg)

}
