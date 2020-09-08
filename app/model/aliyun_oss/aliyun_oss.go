package aliyun_oss

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type AliyunOss struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	config.AliyunOss
	Desc string `gorm:"column:desc" json:"desc"  form:"desc"`
	Status  int8   `gorm:"column:status" json:"status" form:"status" `
	db.TimeStampModel
}

func (m *AliyunOss) TableName() string  {
	return "aliyun_oss"
}
//添加
func (m *AliyunOss) Add(AliyunOss AliyunOss) (ID int,err error) {
	err=db.MysqlConn.Create(&AliyunOss).Error
	ID = AliyunOss.ID
	return
}
//修改
func (m *AliyunOss) Update(data AliyunOss ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&AliyunOss{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *AliyunOss) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&AliyunOss{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *AliyunOss) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&AliyunOss{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&AliyunOss{}).Error
	return
}
func (m *AliyunOss) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *AliyunOss) GetOne (request request.IDStatusRequest ) ( mb *AliyunOss,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






