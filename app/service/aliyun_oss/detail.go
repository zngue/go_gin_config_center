package aliyun_oss
import (
	"github.com/zngue/go_gin_config_center/app/model/aliyun_oss"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *aliyun_oss.AliyunOss,err error)  {
	return s.AliyunOss().GetOne(request)
}
