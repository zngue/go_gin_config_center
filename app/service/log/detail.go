package log
import (
	"github.com/zngue/go_gin_config_center/app/model/log"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *log.Log,err error)  {
	return s.Log().GetOne(request)
}
