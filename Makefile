.DEFAULT_GOAL=all

clean:
	rm -rf bin/

build:
	mkdir -p bin
	go mod tidy
	go mod download
	go build -o bin/kodama ./cmd/kodama

all: clean build

run: all
	./bin/kodama --config=config.yml

air:
	air -c .air.toml 
