package default_load
import "github.com/zngue/go_gin_config_center/app/model/default_load"
func (s *Service) Edit(c default_load.DefaultLoad) (int ,error) {
	if c.ID>0 {
		return s.DefaultLoad().Update(c)
	}else {
		return s.DefaultLoad().Add(c)
	}
}
