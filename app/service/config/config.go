package config
import "github.com/zngue/go_gin_config_center/app/model/config"
type Service struct {
}
func (Service)Config() *config.Config  {
	return new(config.Config)
}
