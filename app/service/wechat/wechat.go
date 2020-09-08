package wechat
import "github.com/zngue/go_gin_config_center/app/model/wechat"
type Service struct {
}
func (Service)Wechat() *wechat.Wechat  {
	return new(wechat.Wechat)
}
