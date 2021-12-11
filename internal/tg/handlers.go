package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimish/client-bot/internal/db"
	"github.com/vladimish/client-bot/pkg/log"
	"github.com/vladimish/client-bot/pkg/utils"
)

func (b *Bot) handleMessage(m *tgbotapi.Message) error {
	state, _, err := db.GetDB().GetUserState(m.Chat.ID)
	if err != nil {
		return err
	}
	switch state {
	case "menu":
		switch m.Text {
		case "Забронировать столик 📋":
			err := b.SendTablesMessage(m.Chat.ID)
			if err != nil {
				return err
			}
			err = db.GetDB().ChangeUserState(m.Chat.ID, "tables", "")
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
		break
	case "tables":
		if m.Text == "⬅" {
			err := b.SendMenuMessage(m.Chat.ID, "Меню")
			if err != nil {
				return err
			}

			err = db.GetDB().ChangeUserState(m.Chat.ID, "menu", "")
			if err != nil {
				return err
			}

			return nil
		}

		tables, err := db.GetDB().GetAllTables()
		if err != nil {
			return err
		}

		tableId := utils.ContainsTable(tables, m.Text)
		if tableId != -1 {
			err = db.GetDB().CreateBookingCallback(m.Chat.ID, tableId)
			if err != nil {
				return err
			}
			err = b.SendTableConfirmationMessage(m.Chat.ID, m.Text)
			if err != nil {
				return err
			}
		} else {
			err := b.SendTextMessage(m.Chat.ID, "Столик не найден")
			if err != nil {
				return err
			}
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

		err = db.GetDB().ChangeUserState(m.Chat.ID, "menu", "")
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) error {
	switch callback.Data {
	case "confirmation":
		_, err := db.GetDB().GetBookingCallback(int64(callback.From.ID))
		if err != nil {
			return err
		}

		log.Get().Info(callback)
		cfg := tgbotapi.NewCallback(callback.ID, "OK")
		_, err = b.bot.AnswerCallbackQuery(cfg)
		if err != nil {
			return err
		}
		break
	default:
		log.Get().Info("Nah")
		break
	}

	return nil
}
