package rules

import (
	"alterGateway/models"
	"alterGateway/pkg"
	"alterGateway/pkg/app"
	"alterGateway/pkg/app/msg"
	"github.com/gin-gonic/gin"
)

func RulesInfo(c *gin.Context)  {
	var Rules models.Rule
	result,err := Rules.GetList()
	pkg.HasError(err, msg.NotFound, 404)
	app.OK(c,result,msg.GetSuccess)
}
