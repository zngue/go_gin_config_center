package we_chat
import "github.com/zngue/go_gin_config_center/app/model/we_chat"
type Service struct {
}
func (Service)WeChat() *we_chat.WeChat  {
	return new(we_chat.WeChat)
}
