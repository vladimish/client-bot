package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vladimish/client-bot/internal/db"
	"github.com/vladimish/client-bot/pkg/log"
)

func (b *Bot) SendTextMessage(chatid int64, text string) error {
	msg := tgbotapi.NewMessage(chatid, text)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) SendMenuMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = buildMenuKeyboard()
	res, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	log.Get().Info("Sent response: ", res.MessageID)

	return nil
}

func (b *Bot) SendTablesMessage(chatID int64) error {
	tables, err := db.GetDB().GetAllTables()
	if err != nil {
		return err
	}

	buttons := make([][]tgbotapi.KeyboardButton, len(tables)/2)
	k := 0
	for i := range buttons {
		buttons[i] = make([]tgbotapi.KeyboardButton, 2)
		buttons[i][0] = tgbotapi.NewKeyboardButton(tables[k].Name)
		k++
		buttons[i][1] = tgbotapi.NewKeyboardButton(tables[k].Name)
		k++
	}

	if len(tables)%2 == 1 {
		buttons = append(buttons, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton(tables[k].Name)})
	}

	buttons = append(buttons, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton("⬅")})

	keyboard := tgbotapi.NewReplyKeyboard(buttons...)
	msg := tgbotapi.NewMessage(chatID, "Список столиков")
	msg.ReplyMarkup = keyboard

	res, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	log.Get().Info("Message sent ", res.MessageID)
	return nil
}

func (b *Bot) SendTableConfirmationMessage(chatId int64, tableName string) error {
	msg := tgbotapi.NewMessage(chatId, tableName)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить бронирование", "confirmation"),
		),
	)

	msg.ReplyMarkup = keyboard
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
