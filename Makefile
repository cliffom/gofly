BINARY_NAME=gofly

dep:
	go get

build:
	go build -o ${BINARY_NAME} *.go

debug:
	go build -gcflags "all=-N -l" -o ${BINARY_NAME} *.go
	./gofly -flies 1 -frametime 100

clean:
	rm ${BINARY_NAME}

run:
	go run main.go fly.go backyard.go critter.go color.go