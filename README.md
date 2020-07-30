# line_bot_weather_forecast
天気予報LINE BOT

## 起動方法

```
$ make up
```

## 停止方法

```
$ make down
```

## Usage
dockerを起動している状態で下記コマンドを入力
```
$ curl -s http://0.0.0.0:8000/api | jq .
```