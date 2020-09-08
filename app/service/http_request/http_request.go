package http_request
import "github.com/zngue/go_gin_config_center/app/model/http_request"
type Service struct {
}
func (Service)HttpRequest() *http_request.HttpRequest  {
	return new(http_request.HttpRequest)
}
