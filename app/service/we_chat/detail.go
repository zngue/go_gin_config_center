package we_chat
import (
	"github.com/zngue/go_gin_config_center/app/model/we_chat"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *we_chat.WeChat,err error)  {
	return s.WeChat().GetOne(request)
}
