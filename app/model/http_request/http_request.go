package http_request

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type HttpRequest struct {
	ID          int    `gorm:"auto_increment" json:"id" form:"id"`
	Desc        string `gorm:"column:desc" json:"desc" form:"desc" `
	Status      int8   `gorm:"column:status" json:"status" form:"status" `
	config.HttpRequest
	db.TimeStampModel
}

func (m *HttpRequest) TableName() string  {
	return "http_request"
}
//添加
func (m *HttpRequest) Add(HttpRequest HttpRequest) (ID int,err error) {
	err=db.MysqlConn.Create(&HttpRequest).Error
	ID = HttpRequest.ID
	return
}
//修改
func (m *HttpRequest) Update(data HttpRequest ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&HttpRequest{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *HttpRequest) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&HttpRequest{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *HttpRequest) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&HttpRequest{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&HttpRequest{}).Error
	return
}
func (m *HttpRequest) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *HttpRequest) GetOne (request request.IDStatusRequest ) ( mb *HttpRequest,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






