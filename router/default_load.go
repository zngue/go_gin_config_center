package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/default_load"
)
func DefaultLoadRouter(g *gin.RouterGroup)  {
	api:=g.Group("default_load")
	{
		api.GET("list",default_load.List)
		api.POST("del_status",default_load.DelOrStatus)
		api.GET("detail",default_load.Detail)
		api.POST("edit",default_load.Edit)
	}
}