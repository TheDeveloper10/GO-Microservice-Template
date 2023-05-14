get_dependencies:
	go get -v ./...

test:
	go test -v ./...

test_race:
	go test -v -race ./...

build:
	go build -o bin/main main.go

run:
	go run main.go

build_docker_image:
	sudo docker build .

build_and_tag_docker_image:
	sudo docker build -t go-microservice-template:1.0.0 .