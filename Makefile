build:
	go build -o mirage.exe main.go

run:
	./mirage.exe serve mirage.json


example:
	./mirage.exe serve --example

start: build run

start-example: build example

