package main

import (
	"database/sql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

// TODO: read token from envvar

func botInit(db *sql.DB) {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TBOT_API_1"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		handleUpd(&update, bot, db)
	}
}

var TMP_ID int = 0

func handleUpd(update *tgbotapi.Update, bot *tgbotapi.BotAPI, db *sql.DB) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	var args = update.Message.CommandArguments()
	if args != "" {
		var cmd = update.Message.Command()
		if cmd == "np" || cmd == "new post" || cmd == "newp" {
			TMP_ID = addPostName(db, args)
			var msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter post contents")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	} else {
		var content = update.Message.Text
		if TMP_ID != 0 && content != "" {
			addPostContent(db, TMP_ID, content)
			TMP_ID = 0
		}
	}
}
