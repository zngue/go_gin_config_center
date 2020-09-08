package service_list_type
import "github.com/zngue/go_gin_config_center/app/model/service_list_type"
func (s *Service) Edit(c service_list_type.ServiceListType) (int ,error) {
	if c.ID>0 {
		return s.ServiceListType().Update(c)
	}else {
		return s.ServiceListType().Add(c)
	}
}
