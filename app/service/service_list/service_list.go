package service_list
import "github.com/zngue/go_gin_config_center/app/model/service_list"
type Service struct {
}
func (Service)ServiceList() *service_list.ServiceList  {
	return new(service_list.ServiceList)
}
