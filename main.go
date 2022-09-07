package main

import (
	logger "github.com/sirupsen/logrus"

	"xxqg-tk/internal/controller"
	"xxqg-tk/internal/initial"
	"xxqg-tk/internal/model"
	"xxqg-tk/pkg/config"
	"xxqg-tk/pkg/database"
	"xxqg-tk/pkg/server"
)

func main() {

	database.DB.Sync2(model.SyncModel...)

	initial.InitData()

	controller.InitRouter()
	logger.Error(server.Start(":" + config.GetString("server.port")))
}
