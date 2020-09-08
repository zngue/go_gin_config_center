package we_chat
import "github.com/zngue/go_gin_config_center/app/model/we_chat"
func (s *Service) Edit(c we_chat.WeChat) (int ,error) {
	if c.ID>0 {
		return s.WeChat().Update(c)
	}else {
		return s.WeChat().Add(c)
	}
}
