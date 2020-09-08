package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/wechat"
)
func WechatRouter(g *gin.RouterGroup)  {
	api:=g.Group("wechat")
	{
		api.GET("list",wechat.List)
		api.POST("del_status",wechat.DelOrStatus)
		api.GET("detail",wechat.Detail)
		api.POST("edit",wechat.Edit)
	}
}