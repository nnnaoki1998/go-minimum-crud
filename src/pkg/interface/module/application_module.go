package module

import (
	"go-minimum-crud/src/pkg/application/service"
)

type ApplicationModule struct {
	UserService service.UserService
}

func InitApplication(daoModule DaoModule) (*ApplicationModule, error) {
	userService := service.UserService{UserRepository: &daoModule.UserDao}
	return &ApplicationModule{userService}, nil
}
