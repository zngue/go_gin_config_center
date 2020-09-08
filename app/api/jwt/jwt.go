package jwt
import "github.com/zngue/go_gin_config_center/app/service/jwt"
func Service() *jwt.Service  {
	return new(jwt.Service)
}
