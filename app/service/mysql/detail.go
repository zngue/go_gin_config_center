package mysql
import (
	"github.com/zngue/go_gin_config_center/app/model/mysql"
	"github.com/zngue/go_tool/src/common/request"
)
func (s *Service) GetOne(request request.IDStatusRequest) (mb *mysql.Mysql,err error)  {
	return s.Mysql().GetOne(request)
}
