package router
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/api/mysql"
)
func MysqlRouter(g *gin.RouterGroup)  {
	api:=g.Group("mysql")
	{
		api.GET("list",mysql.List)
		api.POST("del_status",mysql.DelOrStatus)
		api.GET("detail",mysql.Detail)
		api.POST("edit",mysql.Edit)
	}
}