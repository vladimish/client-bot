package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimish/client-bot/internal/cfg"
	"github.com/vladimish/client-bot/internal/tg"
	"log"
)

func main() {
	// Authorize bot.
	api, err := Authorize()
	if err != nil {
		log.Fatal(err)
	}

	bot := tg.NewBot(api)
	bot.Start()
}

func Authorize() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.GetConfig().ApiKey)
	if err != nil {
		return bot, err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, nil
}
