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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}
