.PHONY: clean build_docker deploy

build: clean
	export GO111MODULE=on
	env GOOS=linux go build -o main.go -o build/nats-client

build_docker:
	export GO111MODULE=on
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main.go -o build/nats-client
	docker build -t nats-client .

clean:
	rm -rf ./build

start_docker: clean build_docker
	docker run --name=nats-client --env-file=.env nats-client

start_compose: clean
	docker-compose -up --build

start_local: clean build
	build/nats-client
