package service_list_type

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type ServiceListType struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Name 	string `gorm:"column:name" json:"name" form:"name" `
	Desc 	string `gorm:"column:desc" json:"desc" form:"desc" `
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	db.TimeStampModel
}

func (m *ServiceListType) TableName() string  {
	return "service_list_type"
}
//添加
func (m *ServiceListType) Add(ServiceListType ServiceListType) (ID int,err error) {
	err=db.MysqlConn.Create(&ServiceListType).Error
	ID = ServiceListType.ID
	return
}
//修改
func (m *ServiceListType) Update(data ServiceListType ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&ServiceListType{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *ServiceListType) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&ServiceListType{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *ServiceListType) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&ServiceListType{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&ServiceListType{}).Error
	return
}
func (m *ServiceListType) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *ServiceListType) GetOne (request request.IDStatusRequest ) ( mb *ServiceListType,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






