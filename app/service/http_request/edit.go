package http_request
import "github.com/zngue/go_gin_config_center/app/model/http_request"
func (s *Service) Edit(c http_request.HttpRequest) (int ,error) {
	if c.ID>0 {
		return s.HttpRequest().Update(c)
	}else {
		return s.HttpRequest().Add(c)
	}
}
