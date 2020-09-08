package service_list_type
import "github.com/zngue/go_gin_config_center/app/model/service_list_type"
type Service struct {
}
func (Service)ServiceListType() *service_list_type.ServiceListType  {
	return new(service_list_type.ServiceListType)
}
