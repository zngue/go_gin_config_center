package redis
import "github.com/zngue/go_gin_config_center/app/model/redis"
func (s *Service) Edit(c redis.Redis) (int ,error) {
	if c.ID>0 {
		return s.Redis().Update(c)
	}else {
		return s.Redis().Add(c)
	}
}
