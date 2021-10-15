build-publisher:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/publisher cmd/publisher/main.go

build-receiver:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/publisher cmd/receiver/main.go


build: build-publisher build-receiver


deploy-local: build
	sls deploy --stage local --verbose