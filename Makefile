build:
	go build -o mirage.exe main.go

run:
	./mirage.exe serve mirage.json

run-multiple:
	./mirage.exe serve --ports=8081,8082,8083

example:
	./mirage.exe serve --example

guide-en:
	./mirage.exe guide-en

guide-fr:
	./mirage.exe guide-fr

start: build run

start-multiple: build run-multiple

start-example: build example

start-guide-en: build guide-en

start-guide-fr: build guide-fr

