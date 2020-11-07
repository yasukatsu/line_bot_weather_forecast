FROM golang:1.13

RUN mkdir -p /go/src/y_ara/line_bot_weather_forecast

WORKDIR /go/src/y_ara/line_bot_weather_forecast

COPY . .

RUN go get -u github.com/codegangsta/gin
