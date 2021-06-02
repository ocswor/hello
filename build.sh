set GOOS=windows
set GOARCH=amd64
go build -o hello.exe

SET GOOS=linux
SET GOARCH=adm64
go build -o hello
