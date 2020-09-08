package mysql
import "github.com/zngue/go_gin_config_center/app/model/mysql"
func (s *Service) Edit(c mysql.Mysql) (int ,error) {
	if c.ID>0 {
		return s.Mysql().Update(c)
	}else {
		return s.Mysql().Add(c)
	}
}
