package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/config"
)
func ConfigRouter(g *gin.RouterGroup)  {
	api:=g.Group("config")
	{
		api.GET("list",config.List)
		api.POST("del_status",config.DelOrStatus)
		api.GET("detail",config.Detail)
		api.POST("edit",config.Edit)
		api.GET("relation",config.Relation)
	}
}