package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

// LineBot ...
func LineBot(w http.ResponseWriter, r *http.Request) {

	NowTemp(w, r)

	lineBot()

}

func lineBot() {
	handler, err := httphandler.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
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

		now := new(Now)
		v := fmt.Sprintf("今の東京の温度は%v度です。\n", int(now.Temp))
		fmt.Printf("v: %v\n", v)

		for _, event := range events {
			if event.Type != linebot.EventTypeMessage {
				log.Printf("mismatch: event.Type: %v linebot.EventTypeMessage: %v\n", event.Type, linebot.EventTypeMessage)
				return
			}

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Printf("success [message: %v, v: %v\n", message, v)
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(v)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})
}
