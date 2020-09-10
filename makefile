build:
	go build -o octo-cli main.go

release:
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/octo-cli-windows
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/octo-cli-linux
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/octo-cli-darwin
