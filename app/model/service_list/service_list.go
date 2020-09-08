package service_list

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type ServiceList struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	ServiceTypeID int  `gorm:"column:service_type_id" json:"service_type_id" form:"service_type_id" `
	Desc 	string `gorm:"column:desc" json:"desc" form:"desc" `
	config.ServiceList
	db.TimeStampModel
}

func (m *ServiceList) TableName() string  {
	return "service_list"
}
//添加
func (m *ServiceList) Add(ServiceList ServiceList) (ID int,err error) {
	err=db.MysqlConn.Create(&ServiceList).Error
	ID = ServiceList.ID
	return
}
//修改
func (m *ServiceList) Update(data ServiceList ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&ServiceList{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *ServiceList) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&ServiceList{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *ServiceList) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&ServiceList{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&ServiceList{}).Error
	return
}
func (m *ServiceList) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *ServiceList) GetOne (request request.IDStatusRequest ) ( mb *ServiceList,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






