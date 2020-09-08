package service_list
import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/model/service_list"
	"github.com/zngue/go_tool/src/common/response"
)
//编辑
func Edit(ctx *gin.Context)  {
	var mb service_list.ServiceList
	if err:=ctx.ShouldBind(&mb); err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	if id,err:=Service().Edit(mb); err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}else {
		response.HttpOk(ctx,id)
		return
	}
}
