package default_load
import (
	"github.com/zngue/go_gin_config_center/app/model/default_load"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *default_load.DefaultLoad,err error)  {
	return s.DefaultLoad().GetOne(request)
}
