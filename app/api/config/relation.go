package config

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_config_center/app/model/config"
)

func Relation(ctx *gin.Context)  {

	var c config.Config
	c.RelationList(1)

}
