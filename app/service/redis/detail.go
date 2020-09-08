package redis
import (
	"github.com/zngue/go_gin_config_center/app/model/redis"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *redis.Redis,err error)  {
	return s.Redis().GetOne(request)
}
