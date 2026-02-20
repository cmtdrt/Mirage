build:
	go build -o mirage.exe main.go

run:
	./mirage.exe serve mirage.json


example:
	./mirage.exe serve --example

guide-en:
	./mirage.exe guide-en

guide-fr:
	./mirage.exe guide-fr

start: build run

start-example: build example

start-guide-en: build guide-en

start-guide-fr: build guide-fr

