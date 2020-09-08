package log
import "github.com/zngue/go_gin_config_center/app/model/log"
func (s *Service) Edit(c log.Log) (int ,error) {
	if c.ID>0 {
		return s.Log().Update(c)
	}else {
		return s.Log().Add(c)
	}
}
