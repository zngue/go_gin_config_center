package service_list
import "github.com/zngue/go_gin_config_center/app/model/service_list"
func (s *Service) Edit(c service_list.ServiceList) (int ,error) {
	if c.ID>0 {
		return s.ServiceList().Update(c)
	}else {
		return s.ServiceList().Add(c)
	}
}
