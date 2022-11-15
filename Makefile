
clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ping functions/ping/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/blog-get functions/blog-get/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/blog-create functions/blog-create/main.go

start: build
	sls offline --useDocker start  --host 0.0.0.0

zip: build
	@zip -j -9 ./bin/blog-get.zip ./bin/blog-get
	@zip -j -9 ./bin/blog-create.zip ./bin/blog-create
	@zip -j -9 ./bin/ping.zip ./bin/ping

format:
	gofmt -s -w .