.PHONY:run
run:
	go run main.go

.PHONY:build
build:
	docker build . -t hey --build-arg HOST_ARCH=amd64 --platform=linux/amd64

.PHONY:fmt
fmt:
	go fmt ./...