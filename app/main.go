package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func main() {
	handler, err := httphandler.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		log.Println(bot,events)
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				log.Println(event)
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	http.Handle("/callback", handler)

	http.Handle("/health",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	http.Handle("/send",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var messages []linebot.SendingMessage

		msg := &linebot.TextMessage{
			Text:    "hogehoge",
			Emojis:  []*linebot.Emoji{},
			Mention: &linebot.Mention{},
		}
		messages = append(messages, msg)

		// append some message to messages

		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		_, err = bot.PushMessage("xxxx", messages...).Do()
		if err != nil {
			// Do something when some bad happened
		}
	}))
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
