package main

import (
	"encoding/json"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func customLogger(logger ...*log.Logger) tele.MiddlewareFunc {
	var l *log.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			data, _ := json.Marshal(c.Update())
			l.Println(string(data))
			return next(c)
		}
	}
}

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
	bot.Use(customLogger())
	setUpHandlers(bot)
	bot.Start()
}
