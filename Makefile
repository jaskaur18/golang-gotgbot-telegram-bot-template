
APP_NAME=golang-tg-bot

dev:
	air -c .air.toml
build:
	CGO_ENABLED=1 go build -o bin/$(APP_NAME) main.go
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)_li main.go
run:
	./bin/$(APP_NAME)
run-linux:
	./bin/$(APP_NAME)_li
