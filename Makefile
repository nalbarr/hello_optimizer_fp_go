PROJ=hello_optimizer_fp_go
SRC=main.go
BIN=hello_optimizer_fp_go

all: run

clean:
	rm $(BIN)

build: $(SRC)
	go build -o $(BIN)

run: build
	go run $(BIN)
