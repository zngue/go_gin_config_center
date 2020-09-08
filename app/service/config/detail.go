package config
import (
	"github.com/zngue/go_gin_config_center/app/model/config"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *config.Config,err error)  {
	return s.Config().GetOne(request)
}
