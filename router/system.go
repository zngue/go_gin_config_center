package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/system"
)
func SystemRouter(g *gin.RouterGroup)  {
	api:=g.Group("system")
	{
		api.GET("list",system.List)
		api.POST("del_status",system.DelOrStatus)
		api.GET("detail",system.Detail)
		api.POST("edit",system.Edit)
	}
}