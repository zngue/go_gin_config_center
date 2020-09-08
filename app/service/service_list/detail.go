package service_list
import (
	"github.com/zngue/go_gin_config_center/app/model/service_list"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *service_list.ServiceList,err error)  {
	return s.ServiceList().GetOne(request)
}
