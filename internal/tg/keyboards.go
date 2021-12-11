package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func buildMenuKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	var buttons [][]tgbotapi.KeyboardButton
	row1 := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Забронировать столик 📋"),
	)

	buttons = append(buttons, row1)
	keyboard := tgbotapi.NewReplyKeyboard(buttons...)

	return &keyboard
}

func buildBackKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	var buttons [][]tgbotapi.KeyboardButton
	row1 := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⬅"),
	)

	buttons = append(buttons, row1)
	keyboard := tgbotapi.NewReplyKeyboard(buttons...)

	return &keyboard
}
