package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/redis"
)
func RedisRouter(g *gin.RouterGroup)  {
	api:=g.Group("redis")
	{
		api.GET("list",redis.List)
		api.POST("del_status",redis.DelOrStatus)
		api.GET("detail",redis.Detail)
		api.POST("edit",redis.Edit)
	}
}