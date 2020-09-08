package wechat

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Wechat struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	db.TimeStampModel
}

func (m *Wechat) TableName() string  {
	return "wechat"
}
//添加
func (m *Wechat) Add(Wechat Wechat) (ID int,err error) {
	err=db.MysqlConn.Create(&Wechat).Error
	ID = Wechat.ID
	return
}
//修改
func (m *Wechat) Update(data Wechat ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Wechat{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Wechat) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Wechat{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Wechat) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Wechat{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Wechat{}).Error
	return
}
func (m *Wechat) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Wechat) GetOne (request request.IDStatusRequest ) ( mb *Wechat,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






