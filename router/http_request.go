package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/http_request"
)
func HttpRequestRouter(g *gin.RouterGroup)  {
	api:=g.Group("http_request")
	{
		api.GET("list",http_request.List)
		api.POST("del_status",http_request.DelOrStatus)
		api.GET("detail",http_request.Detail)
		api.POST("edit",http_request.Edit)
	}
}