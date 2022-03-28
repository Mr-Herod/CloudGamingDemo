set GOARCH=amd64
set GOOS=linux
cd ./Gateway
go mod tidy
go build -o serverGateway
cd ../Naming
go mod tidy
go build -o serverNaming
cd ../Gaming
go mod tidy
go build -o serverGaming
cd ../Account
go mod tidy
go build -o serverAccount
cd ../Record
go mod tidy
go build -o serverRecord
cd ..
echo build process done