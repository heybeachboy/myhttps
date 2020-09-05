
all: server client autotest

server:
	go build -o ./bin/server ./cmd/server

client:
	go build -o ./bin/client ./cmd/client

autotest:
	go build -o ./bin/autotest ./cmd/autotest
