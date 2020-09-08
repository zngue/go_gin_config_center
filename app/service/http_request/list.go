package http_request

import (
	"github.com/zngue/go_gin_config_center/app/model/http_request"
	req "github.com/zngue/go_gin_config_center/app/request/http_request"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*http_request.HttpRequest,err error ) {
	dbConn:=s.HttpRequest().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	request.ListInit()
    if request.DBConn!=nil{
        dbConn = request.DBConn(dbConn)
    }
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&http_request.HttpRequest{}).Count(&num).Error
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