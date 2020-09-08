package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/service_list_type"
)
func ServiceListTypeRouter(g *gin.RouterGroup)  {
	api:=g.Group("service_list_type")
	{
		api.GET("list",service_list_type.List)
		api.POST("del_status",service_list_type.DelOrStatus)
		api.GET("detail",service_list_type.Detail)
		api.POST("edit",service_list_type.Edit)
	}
}