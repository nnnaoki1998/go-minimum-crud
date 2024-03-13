package main

import (
	"fmt"
	"go-minimum-crud/src/pkg/config"
	"go-minimum-crud/src/pkg/interface/module"
	"go-minimum-crud/src/pkg/interface/route"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.LoadConf("conf/server.yaml")
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	daoModule, err := module.InitDao(conf.MysqlConf)
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	applicationModule, err := module.InitApplication(*daoModule)
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}

	userRoute := route.UserRoute{UserService: applicationModule.UserService}

	userRoute.SetRouter()

	port := fmt.Sprintf(":%d", conf.HttpConf.Port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
}
