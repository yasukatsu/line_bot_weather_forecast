package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBot ...
func LineBot(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		linebot.WithHTTPClient(client),
	)
	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Fatalf("failed bot.ParseRequest: %+v", err)
		return
	}

	for _, event := range events {
		if event.Type != linebot.EventTypeMessage {
			log.Printf("mismatch: event.Type: %v linebot.EventTypeMessage: %v\n", event.Type, linebot.EventTypeMessage)
			return
		}

		if event.Type == linebot.EventTypeMessage {
			returnQuickReply(bot, event)
		}
		// switch message := event.Message.(type) {
		// case *linebot.TextMessage:
		// 	// returnTemp(bot, event, message)
		// 	returnQuickReply(bot, event)
		// }
	}
}

func returnTemp(bot *linebot.Client, event *linebot.Event, message linebot.Message) {
	v := fmt.Sprintf("今の東京の温度は%v度です。\n", GetTemp())
	log.Printf("success [message: %v], send: %v\n", message, v)
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(v)).Do(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func returnQuickReply(bot *linebot.Client, event *linebot.Event) {

	var messages []linebot.SendingMessage

	// leftBtn := linebot.NewMessageAction("left", "left clicked")
	// rightBtn := linebot.NewMessageAction("right", "right clicked")
	// template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
	// message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

	message := linebot.NewTextMessage("一覧から選んでね").WithQuickReplies(
		linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton("", linebot.NewMessageAction("スプラッシュボム", "ジョナサン")),
			linebot.NewQuickReplyButton("", linebot.NewMessageAction("ライトニングフィスト", "ボブ")),
		),
	)

	messages = append(messages, message)

	if _, err := bot.ReplyMessage(event.ReplyToken, messages...).Do(); err != nil {
		log.Fatalf("%+v", err)
	}
}
