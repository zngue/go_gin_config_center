package default_load
import "github.com/zngue/go_gin_config_center/app/model/default_load"
type Service struct {
}
func (Service)DefaultLoad() *default_load.DefaultLoad  {
	return new(default_load.DefaultLoad)
}
