package config
import "github.com/zngue/go_gin_config_center/app/service/config"
func Service() *config.Service  {
	return new(config.Service)
}
