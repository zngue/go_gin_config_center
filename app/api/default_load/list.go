package default_load
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/request/default_load"
	"github.com/zngue/go_tool/src/common/response"
)
func List(ctx *gin.Context)  {
	var req  default_load.ListRequest
	if err:=ctx.ShouldBind(&req); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	num,list,errs :=Service().List(req)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,gin.H{
		"count":num,
		"item":list,
	})
	return
}
