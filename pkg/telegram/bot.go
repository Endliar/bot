/*
Здесь мы создадим собственную структуру, она будет нам предоставлять
публичные методы, которые мы сможем вызывать уже в нашем main.go файле и здесь
уже будет реализована сама логика под капотом нашего бота т.е мы создадим обёртку
над нашим основным ботом
*/

package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI // У нас будет поле ссылка на объект, который мы создаём с помощью NewBotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot { // конструктор, который принимает ссылку на наш объект бот
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0) // создаём новую конфигурацию для получения обновлений
	u.Timeout = 60

	updates, err := b.initUpdatesChannel() // создаём канал, в который мы будем получать значения от api
	if err != nil {
		return err
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates { // структура обновлений
		if update.Message == nil { // If we got a message
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0) // создаём новую конфигурацию для получения обновлений
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u) // создаём канал, в который мы будем получать значения от api
}
