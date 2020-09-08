package mysql
import "github.com/zngue/go_gin_config_center/app/service/mysql"
func Service() *mysql.Service  {
	return new(mysql.Service)
}
