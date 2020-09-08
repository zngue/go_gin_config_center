package router
import (
	"github.com/gin-gonic/gin"
)
func Router(g *gin.RouterGroup)  {
	// 
	WeChatRouter(g) 
	WechatRouter(g) 
	ServiceListTypeRouter(g) 
	ConfigRouter(g) 
	ServiceListRouter(g) 
	AliyunOssRouter(g) 
	HttpRequestRouter(g) 
	SystemRouter(g) 
	LogRouter(g) 
	RedisRouter(g) 
	DefaultLoadRouter(g) 
	JwtRouter(g) 
	MysqlRouter(g)
}
