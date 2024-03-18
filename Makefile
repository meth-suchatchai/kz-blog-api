CONFIG_FILE := config.toml

APPLICATION_NAME := $(shell toml get $(CONFIG_FILE) server.application_name)

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o kz_api-linux-x64 ./bootstrap/*.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o kz_api-darwin-x64 ./bootstrap/*.go
upload:
	scp -r ./kz_api-linux-x64 root@178.128.49.82:/root/application/kz-blog-api/
	scp -r ./kz_api-darwin-x64 root@178.128.49.82:/root/application/kz-blog-api/
	scp -r ./views root@178.128.49.82:/root/application/kz-blog-api/
	scp -r ./public root@178.128.49.82:/root/application/kz-blog-api/
test:
	echo $(APPLICATION_NAME)