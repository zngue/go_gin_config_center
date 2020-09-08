package jwt
import (
	"github.com/zngue/go_gin_config_center/app/model/jwt"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *jwt.Jwt,err error)  {
	return s.Jwt().GetOne(request)
}
