.PHONY:run
run:
	go run main.go

.PHONY:build-m1
build-m1:
	docker build . -t hey --build-arg HOST_ARCH=amd64 --platform=linux/amd64

.PHONY:fmt
fmt:
	go fmt ./...