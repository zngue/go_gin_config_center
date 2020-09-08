package system
import "github.com/zngue/go_gin_config_center/app/model/system"
type Service struct {
}
func (Service)System() *system.System  {
	return new(system.System)
}
