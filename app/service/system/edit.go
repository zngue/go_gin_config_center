package system
import "github.com/zngue/go_gin_config_center/app/model/system"
func (s *Service) Edit(c system.System) (int ,error) {
	if c.ID>0 {
		return s.System().Update(c)
	}else {
		return s.System().Add(c)
	}
}
