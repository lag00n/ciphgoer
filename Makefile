BIN_NAME=main.out
SRC_PATH=./src

all: build


build: 
	go build -o ${BIN_NAME} ${SRC_PATH}/main.go

clean:
	go clean
	rm ${BIN_NAME}
run: build
	./${BIN_NAME} ${args}

deps:
	go get fyne.io/fyne/v2
