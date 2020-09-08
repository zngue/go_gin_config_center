package redis
import "github.com/zngue/go_gin_config_center/app/model/redis"
type Service struct {
}
func (Service)Redis() *redis.Redis  {
	return new(redis.Redis)
}
