.PHONY: run get all clean

run: all
	./mute-them-all

get:
	go get -u "github.com/ChimeraCoder/anaconda"

all:
	gofmt -w .
	go build -o mute-them-all github.com/ledyba/go-twitter-mute-them-all/...


clean:
	go clean github.com/ledyba/go-twitter-mute-them-all/...
