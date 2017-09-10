package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"fmt"
)

func GetHomeKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}

	android := tgbotapi.KeyboardButton{Text: AndroidDownload}
	iOS := tgbotapi.KeyboardButton{Text: iOSDownload}
	iOsHelp := tgbotapi.KeyboardButton{Text: iOSInstallHelp}

	commands := [][]tgbotapi.KeyboardButton{}

	row1 := []tgbotapi.KeyboardButton{android}
	commands = append(commands, row1)

	row2 := []tgbotapi.KeyboardButton{iOS}
	commands = append(commands, row2)
	row3 := []tgbotapi.KeyboardButton{iOsHelp}
	commands = append(commands, row3)

	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}
