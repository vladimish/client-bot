package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimish/client-bot/pkg/log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(api *tgbotapi.BotAPI) *Bot {
	b := &Bot{
		bot: api,
	}
	return b
}

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message.IsCommand() {
			log.Get().Info("Handling command ", update.UpdateID)
			b.handleCommand(update.Message)
		} else if update.Message.Text != "" {
			log.Get().Info("Handling update ", update.UpdateID)
			b.handleMessage(update.Message)
		}

	}
}
