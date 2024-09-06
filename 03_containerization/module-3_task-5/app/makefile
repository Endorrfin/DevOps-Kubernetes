BINARY_NAME=myapp
IMAGE_TAG=myapp:latest

all: linux arm mac windows

linux:
    GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux

arm:
    GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-arm

mac:
    GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-mac

windows:
    GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows.exe

docker-build:
    docker build -t $(IMAGE_TAG) .

docker-run:
    docker run --rm -it $(IMAGE_TAG)

clean:
	rm -f $(BINARY_NAME)-*
	docker rmi $(IMAGE_TAG)