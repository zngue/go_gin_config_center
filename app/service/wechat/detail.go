package wechat
import (
	"github.com/zngue/go_gin_config_center/app/model/wechat"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *wechat.Wechat,err error)  {
	return s.Wechat().GetOne(request)
}
