package we_chat

import (
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) DelOrStatus(request request.IDStatusRequest) error {
	if request.From == 1 {
		return s.WeChat().UpdateStatus(request)
	}else{
		return s.WeChat().Del(request)
	}
}
