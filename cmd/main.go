package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimish/client-bot/internal/cfg"
	"github.com/vladimish/client-bot/internal/server"
	"github.com/vladimish/client-bot/internal/tg"
	"log"
)

func main() {
	go server.StartApi()

	// Authorize bot.
	api, err := Authorize()
	if err != nil {
		panic(err)
	}

	bot := tg.NewBot(api)
	err = bot.Start()
	if err != nil {
		panic(err)
	}
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
