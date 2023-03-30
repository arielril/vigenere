GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
BINARY=vigenere
GOHOME?=~/go

all: clean tidy build

build:
	env GOARCH=arm64 $(GOBUILD) -v -ldflags="-extldflags=-static" -o ${BINARY} dashboard.go


build-linux:
	env GOOS=linux GOARCH=arm64 $(GOBUILD) -v -ldflags="-extldflags=-static" -o ${BINARY} dashboard.go

move-bin-linux: 
	mv ${BINARY} ${GOHOME}/bin/${BINARY}

tidy:
	$(GOMOD) tidy

clean:
	rm -f ${BINARY}

