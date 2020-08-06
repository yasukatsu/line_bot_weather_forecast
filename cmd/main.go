package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"y_ara/line_bot_weather_forecast/handler"
)

func main() {

	fmt.Printf("%v\n", "start")
	r := mux.NewRouter()
	r.HandleFunc("/api", handler.OpenWhetherMap)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
