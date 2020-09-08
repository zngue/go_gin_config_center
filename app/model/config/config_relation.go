package config

import (
	"encoding/json"
	"fmt"
	"github.com/zngue/go_gin_config_center/app/model/aliyun_oss"
	"github.com/zngue/go_gin_config_center/app/model/default_load"
	"github.com/zngue/go_gin_config_center/app/model/http_request"
	"github.com/zngue/go_gin_config_center/app/model/jwt"
	"github.com/zngue/go_gin_config_center/app/model/log"
	"github.com/zngue/go_gin_config_center/app/model/mysql"
	"github.com/zngue/go_gin_config_center/app/model/redis"
	"github.com/zngue/go_gin_config_center/app/model/service_list"
	"github.com/zngue/go_gin_config_center/app/model/system"
	"github.com/zngue/go_gin_config_center/app/model/we_chat"
	"github.com/zngue/go_gin_config_center/app/redis/cache_key"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
)

type Relation struct {
	Config
	AliyunOss aliyun_oss.AliyunOss 			`json:"aliyunoss" gorm:"ForeignKey:OssID;AssociationForeignKey:ID"`
	DefaultLoad default_load.DefaultLoad 	`json:"defaultLoad" gorm:"ForeignKey:DefaultLoadID;AssociationForeignKey:ID"`
	HttpRequest http_request.HttpRequest 	`json:"http_request" gorm:"ForeignKey:HttpRequestID;AssociationForeignKey:ID"`
	JWT jwt.Jwt 							`json:"jwt" gorm:"ForeignKey:JwtID;AssociationForeignKey:ID"`
	Log log.Log 							`json:"log" gorm:"ForeignKey:LogID;AssociationForeignKey:ID"`
	Mysql mysql.Mysql 						`json:"mysql" gorm:"ForeignKey:MysqlID;AssociationForeignKey:ID"`
	Redis redis.Redis 						`json:"redis" gorm:"ForeignKey:RedisID;AssociationForeignKey:ID"`
	ServiceList []service_list.ServiceList  `json:"serviceList" gorm:"ForeignKey:ServiceTypeID;AssociationForeignKey:ServiceListTypeID"`
	System system.System 					`json:"system" gorm:"ForeignKey:SystemID;AssociationForeignKey:ID"`
	WeChat we_chat.WeChat 					`json:"weChat" gorm:"ForeignKey:SystemID;AssociationForeignKey:ID"`
}
func (c *Config) RelationList(id int )  {
	var r Relation
	errs:=db.MysqlConn.Model(&Config{}).Where("id = ?",id).
		Preload("DefaultLoad").
		Preload("HttpRequest").
		Preload("JWT").
		Preload("Log").
		Preload("Mysql").
		Preload("Redis").
		Preload("System").
		Preload("WeChat").
		Preload("ServiceList").
		Preload("AliyunOss").First(&r).Error
	basteSlr,_:=json.Marshal(&r)
	fmt.Println(errs)
	_,err:=db.RedisConn.Set(fmt.Sprintf(cache_key.ConfigCenter,id), basteSlr,-1).Result()
	jsonStr,_:=db.RedisConn.Get(fmt.Sprintf(cache_key.ConfigCenter,id)).Result()
	var conf config.Config
	json.Unmarshal([]byte(jsonStr),&conf)
	fmt.Println(conf)
	fmt.Println(err)
	fmt.Println(jsonStr)
}
