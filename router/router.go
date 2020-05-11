package router

import (
	"alterCenter/apis/rules"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Static("/static", "./static")
	apiv1 := r.Group("/api/v1")
	{

		apiv1.GET("/rules", rules.RulesInfo)

	}
	log.Println("路由加载成功！")
	return r
}
