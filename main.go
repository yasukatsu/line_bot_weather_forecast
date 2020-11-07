package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"y_ara/line_bot_weather_forecast/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("%v\n", "start")
	r := mux.NewRouter()
	r.HandleFunc("/api", handler.OpenWhetherMap)
	r.HandleFunc("/api2", handler.NowTemp)
	r.HandleFunc("/nowtemp", handler.LineBot)

	// // lineBot(temp)
	// h, err := httphandler.New(
	// 	os.Getenv("CHANNEL_SECRET"),
	// 	os.Getenv("CHANNEL_TOKEN"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8000"
	// 	log.Printf("Defaulting to port %s", port)
	// }

	// // Setup HTTP Server for receiving requests from LINE platform
	// h.HandleEvents(func(events []*linebot.Event, r *http.Request) {
	// 	bot, err := h.NewClient()
	// 	if err != nil {
	// 		log.Print(err)
	// 		return
	// 	}

	// 	// now := new(Now)
	// 	// v := fmt.Sprintf("今の東京の温度は%v度です。\n", int(now.Temp))

	// 	temp := handler.GetTemp()
	// 	v := fmt.Sprintf("今の東京の温度は%v度です。\n", temp)

	// 	fmt.Printf("v: %v\n", v)

	// 	for _, event := range events {
	// 		if event.Type != linebot.EventTypeMessage {
	// 			log.Printf("mismatch: event.Type: %v linebot.EventTypeMessage: %v\n", event.Type, linebot.EventTypeMessage)
	// 			return
	// 		}

	// 		switch message := event.Message.(type) {
	// 		case *linebot.TextMessage:
	// 			log.Printf("success [message: %v, v: %v\n", message, v)
	// 			if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(v)).Do(); err != nil {
	// 				log.Print(err)
	// 				return
	// 			}
	// 		}
	// 	}
	// })
	// http.Handle("/nowtemp", h)
	// log.Printf("Listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatalln(err)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}
