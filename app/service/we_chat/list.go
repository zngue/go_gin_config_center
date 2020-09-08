package we_chat

import (
	"github.com/zngue/go_gin_config_center/app/model/we_chat"
	req "github.com/zngue/go_gin_config_center/app/request/we_chat"
)

func (s *Service) List ( request req.ListRequest ) (num int,mb []*we_chat.WeChat,err error ) {
	dbConn:=s.WeChat().GetList()
	if len(request.IDArr)>0 {
		dbConn=dbConn.Where(" id in (?) ",request.IDArr)
	}
	request.ListInit()
    if request.DBConn!=nil{
        dbConn = request.DBConn(dbConn)
    }
	b,err:=request.OnlyCount(func() error {
		return dbConn.Model(&we_chat.WeChat{}).Count(&num).Error
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