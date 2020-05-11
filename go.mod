module alterGateway

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.12
	github.com/spf13/viper v1.7.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

replace alterGateway/api/rules => ./api/rules
