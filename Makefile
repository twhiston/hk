.PHONY: app

app:
	go build -o bin/hk
	env GOOS=linux GOARCH=amd64 go build -o bin/hk-amd64
	env GOOS=linux GOARCH=386 go build -o bin/hk-386

default: app