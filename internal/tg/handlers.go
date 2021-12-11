package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleMessage(m *tgbotapi.Message) error {
	switch m.Text {
	case "Забронировать столик 📋":
		err := b.SendTablesMessage(m.Chat.ID)
		if err != nil {
			return err
		}
		break
	default:
		err := b.SendTextMessage(m.Chat.ID, "Неизвестная команда")
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func (b *Bot) handleCommand(m *tgbotapi.Message) error {
	switch m.Text {
	case "/start":
		err := b.SendMenuMessage(m.Chat.ID, "Hello")
		if err != nil {
			return err
		}
		break
	}

	return nil
}
