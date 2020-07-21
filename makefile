run:
	go run ./uh/main.go

build:
	go build -o build/uh ./uh

install:
	go install ./uh

test:
	go run ./uh/main.go intellij
	go run ./uh/main.go vscode
	go run ./uh/main.go eclipse

all:
	build run install

.PHONY: run build install all