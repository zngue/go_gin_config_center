package wechat
import "github.com/zngue/go_gin_config_center/app/service/wechat"
func Service() *wechat.Service  {
	return new(wechat.Service)
}
