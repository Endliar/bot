package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввёл команду /start"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text) // логгируем от кого пришло сообщение и какое сообщение
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)  // новая структура message
	//msg.ReplyToMessageID = update.Message.MessageID
	b.bot.Send(msg) //ответ на сообщение пользователя типа

}
