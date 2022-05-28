module github.com/Mr-Herod/CloudGamingDemo/Gateway

go 1.18

replace (
	github.com/Mr-Herod/CloudGamingDemo/Account => ../Account
	github.com/Mr-Herod/CloudGamingDemo/Common => ../Common
	github.com/Mr-Herod/CloudGamingDemo/Gaming => ../Gaming
	github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming
	github.com/Mr-Herod/CloudGamingDemo/Record => ../Record
)

require (
	github.com/Mr-Herod/CloudGamingDemo/Account v0.0.0-00010101000000-000000000000
	github.com/Mr-Herod/CloudGamingDemo/Common v0.0.0-00010101000000-000000000000
	github.com/Mr-Herod/CloudGamingDemo/Gaming v0.0.0-00010101000000-000000000000
	github.com/Mr-Herod/CloudGamingDemo/Record v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	google.golang.org/grpc v1.45.0
)
