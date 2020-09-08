package log

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Log struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	db.TimeStampModel
}

func (m *Log) TableName() string  {
	return "log"
}
//添加
func (m *Log) Add(Log Log) (ID int,err error) {
	err=db.MysqlConn.Create(&Log).Error
	ID = Log.ID
	return
}
//修改
func (m *Log) Update(data Log ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Log{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Log) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Log{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Log) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Log{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Log{}).Error
	return
}
func (m *Log) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Log) GetOne (request request.IDStatusRequest ) ( mb *Log,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






