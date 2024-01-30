BIN_NAME=main.out
SRC_PATH=./src

all: build


build: 
	go build -o ${BIN_NAME} ${SRC_PATH}/main.go

clean:
	go clean
	rm ${BIN_NAME}
run:
	./${BIN_NAME} ${args}

rebuild: build run

deps:
	go get fyne.io/fyne/v2
