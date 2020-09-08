package service_list

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.ServiceList().UpdateStatus(request)
	}else{
		return s.ServiceList().Del(request)
	}
}
