package aliyun_oss
import "github.com/zngue/go_gin_config_center/app/model/aliyun_oss"
func (s *Service) Edit(c aliyun_oss.AliyunOss) (int ,error) {
	if c.ID>0 {
		return s.AliyunOss().Update(c)
	}else {
		return s.AliyunOss().Add(c)
	}
}
