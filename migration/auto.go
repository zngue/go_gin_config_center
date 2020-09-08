package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_config_center/app/model/aliyun_oss"
	"github.com/zngue/go_gin_config_center/app/model/config"
	"github.com/zngue/go_gin_config_center/app/model/default_load"
	"github.com/zngue/go_gin_config_center/app/model/http_request"
	"github.com/zngue/go_gin_config_center/app/model/jwt"
	"github.com/zngue/go_gin_config_center/app/model/log"
	"github.com/zngue/go_gin_config_center/app/model/mysql"
	"github.com/zngue/go_gin_config_center/app/model/redis"
	"github.com/zngue/go_gin_config_center/app/model/service_list"
	"github.com/zngue/go_gin_config_center/app/model/service_list_type"
	"github.com/zngue/go_gin_config_center/app/model/system"
	"github.com/zngue/go_gin_config_center/app/model/we_chat"
)

func Auto(db *gorm.DB)  {

	db.AutoMigrate(
		new(aliyun_oss.AliyunOss),
		new(default_load.DefaultLoad),
		new(http_request.HttpRequest),
		new(jwt.Jwt),
		new(log.Log),
		new(mysql.Mysql),
		new(redis.Redis),
		new(service_list.ServiceList),
		new(service_list_type.ServiceListType),
		new(system.System),
		new(config.Config),
		new(we_chat.WeChat),
	)
}
