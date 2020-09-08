package we_chat

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type WeChat struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	db.TimeStampModel
}

func (m *WeChat) TableName() string  {
	return "we_chat"
}
//添加
func (m *WeChat) Add(WeChat WeChat) (ID int,err error) {
	err=db.MysqlConn.Create(&WeChat).Error
	ID = WeChat.ID
	return
}
//修改
func (m *WeChat) Update(data WeChat ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&WeChat{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *WeChat) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&WeChat{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *WeChat) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&WeChat{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&WeChat{}).Error
	return
}
func (m *WeChat) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *WeChat) GetOne (request request.IDStatusRequest ) ( mb *WeChat,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






