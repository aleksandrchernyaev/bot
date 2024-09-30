package main

import (
	"fmt"
	"log"
	"os"

	commands "github.com/aleksandrchernyaev/bot/internal/app/commander"
	"github.com/aleksandrchernyaev/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	token := os.Getenv("TOKEN")
	fmt.Println("you tiken is: " + token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	prouctService := product.NewService()
	commander := commands.NewCommander(bot, prouctService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}

	}
}
