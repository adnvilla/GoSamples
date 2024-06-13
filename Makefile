default: build

# Linux OR Windows
#go build -o /dev/null ./... OR go build -o NUL ./...

build:
	go build -o /dev/null ./... 

build_test:
	go test -c ./...
