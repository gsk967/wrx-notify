package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"sync"
)

// Telegram to send telegram alert interface
type Telegram interface {
	SendTelegramMessage(msgText, botToken string, chatID int64) error
}

type TelegramAlert struct {
	botapi *tgbotapi.BotAPI
	chatId int64
}

// NewTelegramAlerter returns telegram bot reference
func NewTelegramAlerter(botToken string, chatId int64) *TelegramAlert {
	return &TelegramAlert{botapi: GetTelegramBotInstance(botToken), chatId: chatId}
}

var bot *tgbotapi.BotAPI
var once sync.Once
var err error

func GetTelegramBotInstance(botToken string) *tgbotapi.BotAPI {
	log.Printf("Initialized the telegran bot...")
	once.Do(func() {
		bot, err = tgbotapi.NewBotAPI(botToken)
		if err != nil {
			log.Fatalf("Telegram bot initialization failed, Err : %v", err)
		}
		log.Print("Telegram bot initialization succeed")
	})
	return bot
}

// SendTelegramMessage to send alert to telegram bot
func (t *TelegramAlert) SendTelegramMessage(msgText string) error {
	bot := t.botapi
	//bot.Debug = true
	msg := tgbotapi.NewMessage(t.chatId, msgText)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
