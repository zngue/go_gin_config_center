package default_load

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type DefaultLoad struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	Desc string `gorm:"column:desc" json:"desc"  form:"desc"`
	config.DefaultLoad
	db.TimeStampModel
}

func (m *DefaultLoad) TableName() string  {
	return "default_load"
}
//添加
func (m *DefaultLoad) Add(DefaultLoad DefaultLoad) (ID int,err error) {
	err=db.MysqlConn.Create(&DefaultLoad).Error
	ID = DefaultLoad.ID
	return
}
//修改
func (m *DefaultLoad) Update(data DefaultLoad ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&DefaultLoad{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *DefaultLoad) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&DefaultLoad{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *DefaultLoad) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&DefaultLoad{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&DefaultLoad{}).Error
	return
}
func (m *DefaultLoad) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *DefaultLoad) GetOne (request request.IDStatusRequest ) ( mb *DefaultLoad,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






