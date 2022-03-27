module github.com/Mr-Herod/CloudGamingDemo/Common

go 1.16

replace github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming

require (
	github.com/Mr-Herod/CloudGamingDemo/Naming v0.0.0-20220326145348-3d9140ddbe6d
	google.golang.org/grpc v1.45.0
)
