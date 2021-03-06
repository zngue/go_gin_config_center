package service_list_type

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.ServiceListType().UpdateStatus(request)
	}else{
		return s.ServiceListType().Del(request)
	}
}
