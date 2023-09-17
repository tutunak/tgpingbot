package main

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func setUpHandlers(bot *tele.Bot) {
	bot.Handle("/ping", func(context tele.Context) error {
		return context.Reply("pong")
	})

	bot.Handle("ping", func(context tele.Context) error {
		return context.Reply("pong")
	})

}

func main() {
	conf := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	setUpHandlers(bot)
	bot.Start()
}
