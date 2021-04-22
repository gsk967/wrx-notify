

build:
	go build -o ./build/wrx-notify ./cmd/main.go


install:
	go build -o ./build/wrx-notify ./cmd/main.go
	cp ./build/wrx-notify ${HOME}/go/bin/wrx-notify


start:
	wrx-notify --exchanges-list ./exchanges.json --watch 10