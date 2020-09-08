package http_request
import "github.com/zngue/go_gin_config_center/app/service/http_request"
func Service() *http_request.Service  {
	return new(http_request.Service)
}
