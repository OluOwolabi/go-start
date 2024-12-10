APP_NAME=go-start

init:
	@echo "init"

tidy:
	go mod tidy	

build:
	go build -o build/bin/$(APP_NAME)