package log
import "github.com/zngue/go_gin_config_center/app/service/log"
func Service() *log.Service  {
	return new(log.Service)
}
