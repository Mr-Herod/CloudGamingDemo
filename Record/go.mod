module github.com/Mr-Herod/CloudGamingDemo/Record

go 1.16

replace (
	github.com/Mr-Herod/CloudGamingDemo/Common => ../Common
	github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming
	github.com/Mr-Herod/CloudGamingDemo/Record => ../Record
)

require (
	github.com/Mr-Herod/CloudGamingDemo/Common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
