package main

import (
	"bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6277339678:AAG6gZDM5Z0NA4pbnQaa-JiYGuN1guK4ESg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}