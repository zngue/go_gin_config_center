package default_load

import (
	"github.com/zngue/go_gin_config_center/app/model/default_load"
	req "github.com/zngue/go_gin_config_center/app/request/default_load"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*default_load.DefaultLoad,err error ) {
	dbConn:=s.DefaultLoad().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	request.ListInit()
    if request.DBConn!=nil{
        dbConn = request.DBConn(dbConn)
    }
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&default_load.DefaultLoad{}).Count(&num).Error
	})
	if b || err!=nil{
		return
	}
	if request.PageSize>0 && request.Page.Page>0 {
		dbConn = dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	if request.Order!=nil {
        dbConn = request.Order(dbConn)
    }else{
        dbConn = dbConn.Order("id desc")
    }
	err =dbConn.Find(&mb).Error
	if err != nil {
		return
	}
	return
}