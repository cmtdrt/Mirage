build:
	go build -o mirage.exe main.go

run:
	./mirage.exe serve mirage.json

start: build run