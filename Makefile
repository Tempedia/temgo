

all: temgo

temgo:
	CGO_ENABLED=0 go build -o bin/temgo
