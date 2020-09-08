package config
import "github.com/zngue/go_gin_config_center/app/model/config"
func (s *Service) Edit(c config.Config) (int ,error) {
	if c.ID>0 {
		return s.Config().Update(c)
	}else {
		return s.Config().Add(c)
	}
}
