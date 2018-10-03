build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go





#In Windows
	set GOOS=linux
	set GOARCH=amd64
	set GOARCH=amd64 set GOOS=linux go build -ldflags "-s -w" -o .\bin\hello .\hello\main.go
	set GOARCH=amd64 set GOOS=linux go build -ldflags "-s -w" -o .\bin\world .\world\main.go


	go build -o .\bin\hello .\hello\main.go
	go build -o .\bin\world .\world\main.go

	%USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip .\bin\hello



	go build -v -ldflags "-d -s -w" -a -tags netgo -installsuffix netgo -o .\bin\hello .\hello\main.go
	go build -v -ldflags "-d -s -w" -a -tags netgo -installsuffix netgo -o .\bin\world .\world\main.go