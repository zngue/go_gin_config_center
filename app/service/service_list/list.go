package service_list

import (
	"github.com/zngue/go_gin_config_center/app/model/service_list"
	req "github.com/zngue/go_gin_config_center/app/request/service_list"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*service_list.ServiceList,err error ) {
	dbConn:=s.ServiceList().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	request.ListInit()
    if request.DBConn!=nil{
        dbConn = request.DBConn(dbConn)
    }
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&service_list.ServiceList{}).Count(&num).Error
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