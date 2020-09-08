package default_load
import "github.com/zngue/go_gin_config_center/app/service/default_load"
func Service() *default_load.Service  {
	return new(default_load.Service)
}
