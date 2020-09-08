package redis
import "github.com/zngue/go_gin_config_center/app/service/redis"
func Service() *redis.Service  {
	return new(redis.Service)
}
