package http_request
import (
	"github.com/zngue/go_gin_config_center/app/model/http_request"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *http_request.HttpRequest,err error)  {
	return s.HttpRequest().GetOne(request)
}
