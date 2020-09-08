package mysql
import "github.com/zngue/go_gin_config_center/app/model/mysql"
type Service struct {
}
func (Service)Mysql() *mysql.Mysql  {
	return new(mysql.Mysql)
}
