
all: server client

server:
	go build -o ./bin/server ./cmd/server

client:
	go build -o ./bin/client ./cmd/client
