package log
import "github.com/zngue/go_gin_config_center/app/model/log"
type Service struct {
}
func (Service)Log() *log.Log  {
	return new(log.Log)
}
