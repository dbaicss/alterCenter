package rules

import (
	"alterCenter/models"
	"alterCenter/pkg"
	"alterCenter/pkg/app"
	"alterCenter/pkg/app/msg"
	"github.com/gin-gonic/gin"
)

func RulesInfo(c *gin.Context)  {
	var Rules models.Rule
	result,err := Rules.GetList()
	pkg.HasError(err, msg.NotFound, 404)
	app.OK(c,result,msg.GetSuccess)
}
