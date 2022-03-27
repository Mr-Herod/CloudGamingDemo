module github.com/Mr-Herod/CloudGamingDemo/Gateway

go 1.18

replace (
	github.com/Mr-Herod/CloudGamingDemo/Account => ../Account
	github.com/Mr-Herod/CloudGamingDemo/Common => ../Common
	github.com/Mr-Herod/CloudGamingDemo/Gateway => ../Gateway
	github.com/Mr-Herod/CloudGamingDemo/Naming => ../Naming
)

require (
	github.com/Mr-Herod/CloudGamingDemo/Common v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
