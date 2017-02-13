BIN := go-random-text
TMP_DIR       := .tmp
RELEASE_VER   := $(shell git rev-parse --short HEAD)
BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

.phony: clean 

.DEFAULT_GOAL:	

$(BIN): $(SOURCES) 
	$(BUILD_ENV) go build -v -x -o $(BIN) examples/cli/main.go

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

