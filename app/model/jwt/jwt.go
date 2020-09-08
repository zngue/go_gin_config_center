package jwt

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_tool/src/common/request"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
)

type Jwt struct {
	ID         int    `gorm:"auto_increment" json:"id" form:"id"`
	config.Jwt
	Status  int8       `gorm:"column:status" json:"status" form:"status" `
	Desc string `json:"desc" form:"desc" gorm:"column:desc" `
	db.TimeStampModel
}

func (m *Jwt) TableName() string  {
	return "jwt"
}
//添加
func (m *Jwt) Add(Jwt Jwt) (ID int,err error) {
	err=db.MysqlConn.Create(&Jwt).Error
	ID = Jwt.ID
	return
}
//修改
func (m *Jwt) Update(data Jwt ) (ID int ,err error)  {
	dataInfo:=common.StuckToMap(data)
	err=db.MysqlConn.Model(&Jwt{}).Update(dataInfo).Error
	ID = data.ID
	return
}
//修改状态
func (m *Jwt) UpdateStatus (request request.IDStatusRequest) (err error) {
	err = db.MysqlConn.Model(&Jwt{}).Where( "id in (?)",request.IDArr ).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
func (m *Jwt) Del (statusRequest request.IDStatusRequest) (err error) {
	err=db.MysqlConn.Model(&Jwt{}).Where( "id in (?)",statusRequest.IDArr ).Delete(&Jwt{}).Error
	return
}
func (m *Jwt) GetList() *gorm.DB  {

	return db.MysqlConn
}
func (m *Jwt) GetOne (request request.IDStatusRequest ) ( mb *Jwt,err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).First(&m).Error
	mb = m
	return
}






