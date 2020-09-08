package jwt
import "github.com/zngue/go_gin_config_center/app/model/jwt"
type Service struct {
}
func (Service)Jwt() *jwt.Jwt  {
	return new(jwt.Jwt)
}
