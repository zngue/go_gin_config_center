package aliyun_oss
import "github.com/zngue/go_gin_config_center/app/model/aliyun_oss"
type Service struct {
}
func (Service)AliyunOss() *aliyun_oss.AliyunOss  {
	return new(aliyun_oss.AliyunOss)
}
