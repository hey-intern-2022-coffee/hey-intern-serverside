.PHONY:run
run:
	go run main.go

.PHONY:build
build:
	docker build . -t hey --platform=linux/amd64

.PHONY:fmt
fmt:
	go fmt ./...