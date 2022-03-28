set GOARCH=amd64
set GOOS=linux
go mod tidy
go build -o serverRecord
echo build process done