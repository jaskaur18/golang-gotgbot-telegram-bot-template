
APP_NAME=golang-tg-bot

dev:
	air -c .air.toml
build:
	CGO_ENABLED=1 go build -o bin/$(APP_NAME) main.go
run:
	./bin/$(APP_NAME)