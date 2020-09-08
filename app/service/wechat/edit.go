package wechat
import "github.com/zngue/go_gin_config_center/app/model/wechat"
func (s *Service) Edit(c wechat.Wechat) (int ,error) {
	if c.ID>0 {
		return s.Wechat().Update(c)
	}else {
		return s.Wechat().Add(c)
	}
}
