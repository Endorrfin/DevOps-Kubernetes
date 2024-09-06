.PHONY: linux arm macos windows clean

linux:
    GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64
    docker build -t my-app:linux-amd64 . -f Dockerfile
arm:
    GOOS=linux GOARCH=arm go build -o bin/linux-arm
    docker build -t my-app:linux-arm . -f Dockerfile
macos:
    GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64
    docker build -t my-app:macos-amd64 . -f Dockerfile
windows:
    GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64.exe
    docker build -t my-app:windows-amd64 . -f Dockerfile
clean:
    docker rmi my-app:linux-amd64 my-app:linux-arm my-app:macos-amd64 my-app:windows-amd64 
