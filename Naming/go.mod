module github.com/Mr-Herod/CloudGamingDemo/Naming

go 1.16

replace github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming

require (
	github.com/go-redis/redis/v8 v8.11.5
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
