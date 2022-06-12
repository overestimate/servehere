set GOOS=windows
set GOARCH=amd64
go build .
set GOARCH=arm64
go build .
set GOOS=darwin
go build .
set GOARCH=amd64
go build .
set GOOS=linux
go build .
set GOARCH=arm64
go build .
set GOOS=
set GOARCH=