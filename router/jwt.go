package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/jwt"
)
func JwtRouter(g *gin.RouterGroup)  {
	api:=g.Group("jwt")
	{
		api.GET("list",jwt.List)
		api.POST("del_status",jwt.DelOrStatus)
		api.GET("detail",jwt.Detail)
		api.POST("edit",jwt.Edit)
	}
}