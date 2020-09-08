package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/log"
)
func LogRouter(g *gin.RouterGroup)  {
	api:=g.Group("log")
	{
		api.GET("list",log.List)
		api.POST("del_status",log.DelOrStatus)
		api.GET("detail",log.Detail)
		api.POST("edit",log.Edit)
	}
}