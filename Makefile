build:
	docker-compose up --build -d

up:
	docker-compose up -d

down:
	docker-compose down

log:
	docker-compose logs $(V)

ps:
	docker-compose ps