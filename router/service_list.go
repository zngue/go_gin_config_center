package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/service_list"
)
func ServiceListRouter(g *gin.RouterGroup)  {
	api:=g.Group("service_list")
	{
		api.GET("list",service_list.List)
		api.POST("del_status",service_list.DelOrStatus)
		api.GET("detail",service_list.Detail)
		api.POST("edit",service_list.Edit)
	}
}