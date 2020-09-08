package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/aliyun_oss"
)
func AliyunOssRouter(g *gin.RouterGroup)  {
	api:=g.Group("aliyun_oss")
	{
		api.GET("list",aliyun_oss.List)
		api.POST("del_status",aliyun_oss.DelOrStatus)
		api.GET("detail",aliyun_oss.Detail)
		api.POST("edit",aliyun_oss.Edit)
	}
}