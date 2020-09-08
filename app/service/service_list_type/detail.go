package service_list_type
import (
	"github.com/zngue/go_gin_config_center/app/model/service_list_type"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *service_list_type.ServiceListType,err error)  {
	return s.ServiceListType().GetOne(request)
}
