NAME = golearning
GOOS = linux

build:
	GOOS=linux GOARCH=amd64 go build -o ${NAME} .

context: build
	./${NAME} context