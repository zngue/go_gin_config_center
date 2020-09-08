package system
import (
	"github.com/zngue/go_gin_config_center/app/model/system"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *system.System,err error)  {
	return s.System().GetOne(request)
}
