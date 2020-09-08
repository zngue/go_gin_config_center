package system

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type System struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status"  form:"status"  json:"status"`
	Desc string `json:"desc" form:"desc" gorm:"column:desc" `
	config.System
	db.TimeStampModel
}

func (m *System) TableName() string  {
	return "system"
}
//添加
func (m *System) Add(System System) (ID int,err error) {
	err=db.MysqlConn.Create(&System).Error
	ID = System.ID
	return
}
//修改
func (m *System) Update(data System ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&System{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *System) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&System{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *System) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&System{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&System{}).Error
	return
}
func (m *System) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *System) GetOne (request request.IDStatusRequest ) ( mb *System,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






