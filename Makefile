.PHONY: all build run get clean

all: build

run:
	go run main.go

build:
	go build -o startAutoscale

get:
	go mod tidy
	go mod download

clean:
	rm -f startAutoscale
