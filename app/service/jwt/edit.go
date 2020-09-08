package jwt
import "github.com/zngue/go_gin_config_center/app/model/jwt"
func (s *Service) Edit(c jwt.Jwt) (int ,error) {
	if c.ID>0 {
		return s.Jwt().Update(c)
	}else {
		return s.Jwt().Add(c)
	}
}
