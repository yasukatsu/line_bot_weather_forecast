FROM golang:1.13

RUN mkdir -p /go/src/y_ara/line_bot_weather_forecast

WORKDIR /go/src/y_ara/line_bot_weather_forecast

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go get github.com/line/line-bot-sdk-go/linebot/httphandler