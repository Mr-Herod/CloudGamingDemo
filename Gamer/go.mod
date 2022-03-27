module github.com/Mr-Herod/CloudGamingDemo/Gamer

go 1.18

replace (
	github.com/Mr-Herod/CloudGamingDemo/Account => ../Account
	github.com/Mr-Herod/CloudGamingDemo/Common => ../Common
	github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming
)

require (
	github.com/Mr-Herod/CloudGamingDemo/Common v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/pion/webrtc/v3 v3.1.27
)
