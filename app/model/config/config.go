package config

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Config struct {
	ID         int    		`gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       		`gorm:"column:status" json:"status" form:"status" `
	OssID int  				`gorm:"column:oss_id" json:"oss_id" form:"oss_id" `
	DefaultLoadID int  		`gorm:"column:default_load_id" json:"default_load_id" form:"default_load_id" `
	HttpRequestID int  		`gorm:"column:http_request_id" json:"http_request_id" form:"http_request_id" `
	JwtID int  				`gorm:"column:jwt_id" json:"jwt_id" form:"jwt_id" `
	LogID int  				`gorm:"column:log_id" json:"log_id" form:"log_id" `
	MysqlID int  			`gorm:"column:mysql_id" json:"mysql_id" form:"mysql_id" `
	RedisID int  			`gorm:"column:redis_id" json:"redis_id" form:"redis_id" `
	ServiceListTypeID int  	`gorm:"column:service_list_type_id" json:"service_list_type_id" form:"service_list_type_id" `
	SystemID int  			`gorm:"column:system_id" json:"system_id" form:"system_id" `
	WeChatID int  			`gorm:"column:we_chat_id" json:"we_chat_id" form:"we_chat_id" `
	db.TimeStampModel
}
func (c *Config) TableName() string  {
	return "config"
}
//添加
func (c *Config) Add(Config Config) (ID int,err error) {
	err=db.MysqlConn.Create(&Config).Error
	ID = Config.ID
	c.RelationList(ID)
	return
}
//修改
func (c *Config) Update(data Config ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Config{}).Update(dataInfo).Error
	ID = data.ID
	c.RelationList(ID)
	return
}
//修改状态
func (c *Config) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Config{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (c *Config) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Config{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Config{}).Error
	return
}
func (c *Config) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (c *Config) GetOne (request request.IDStatusRequest ) ( mb *Config,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&c).Error
	mb = c
	return
}






