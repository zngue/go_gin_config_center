package redis

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Redis struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status"  form:"status"  json:"status"`
	Desc string `json:"desc" form:"desc" gorm:"column:desc" `
	config.Redis
	db.TimeStampModel
}

func (m *Redis) TableName() string  {
	return "redis"
}
//添加
func (m *Redis) Add(Redis Redis) (ID int,err error) {
	err=db.MysqlConn.Create(&Redis).Error
	ID = Redis.ID
	return
}
//修改
func (m *Redis) Update(data Redis ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Redis{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Redis) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Redis{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Redis) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Redis{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Redis{}).Error
	return
}
func (m *Redis) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Redis) GetOne (request request.IDStatusRequest ) ( mb *Redis,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






