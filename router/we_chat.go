package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/we_chat"
)
func WeChatRouter(g *gin.RouterGroup)  {
	api:=g.Group("we_chat")
	{
		api.GET("list",we_chat.List)
		api.POST("del_status",we_chat.DelOrStatus)
		api.GET("detail",we_chat.Detail)
		api.POST("edit",we_chat.Edit)
	}
}