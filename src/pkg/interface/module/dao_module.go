package module

import (
	"go-minimum-crud/src/pkg/config"
	"go-minimum-crud/src/pkg/infrastructure/dao"
)

type DaoModule struct {
	UserDao dao.UserDao
}

func InitDao(mysqlConf config.MysqlConf) (*DaoModule, error) {
	db, err := dao.CreateDB(
		mysqlConf.User,
		mysqlConf.Passward,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DatabaseName,
	)
	if err != nil {
		return nil, err
	}
	var userDao dao.UserDao = dao.UserDao{DB: db}
	return &DaoModule{userDao}, nil
}
