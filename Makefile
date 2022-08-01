BINARY_NAME=gofly

dep:
	go get

build:
	go build -o ${BINARY_NAME} *.go

clean:
	rm ${BINARY_NAME}

run:
	go run main.go fly.go backyard.go critter.go color.go