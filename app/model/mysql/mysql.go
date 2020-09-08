package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Mysql struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	config.Mysql
	Desc string `json:"desc" form:"desc" gorm:"column:desc" `
	db.TimeStampModel
}

func (m *Mysql) TableName() string  {
	return "mysql"
}
//添加
func (m *Mysql) Add(Mysql Mysql) (ID int,err error) {
	err=db.MysqlConn.Create(&Mysql).Error
	ID = Mysql.ID
	return
}
//修改
func (m *Mysql) Update(data Mysql ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Mysql{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Mysql) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Mysql{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Mysql) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Mysql{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Mysql{}).Error
	return
}
func (m *Mysql) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Mysql) GetOne (request request.IDStatusRequest ) ( mb *Mysql,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






