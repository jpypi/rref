OUT_NAME = rref
OUT_DIR = ./bin
#CROSS_COMPILE = env GOOS=windows|linux|darwin GOARCH=386|arm GOARM=6|7|11


.PHONY: all
all: main.go matrix.go fraction.go
	go build -o ${OUT_DIR}/${OUT_NAME} $^

.PHONY: run
run:
	${OUT_DIR}/${OUT_NAME}

.PHONY: brun
brun: all run
