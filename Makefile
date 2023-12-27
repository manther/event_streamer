BINARY_NAME=streaming_events

VERSION?=0.0.0

build: 
		go build -o bin/app 
run:  
		build ./bin/app  
test: 
		go test -v -cover ./... -count=1