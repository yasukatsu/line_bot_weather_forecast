package main

import (
	"fmt"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))

}
