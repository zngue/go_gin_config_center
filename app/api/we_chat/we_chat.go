package we_chat
import "github.com/zngue/go_gin_config_center/app/service/we_chat"
func Service() *we_chat.Service  {
	return new(we_chat.Service)
}
