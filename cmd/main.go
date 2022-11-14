package main

import (
	"SQLSimulator/controller/api"
	"SQLSimulator/entity"
	"SQLSimulator/logging"
	"SQLSimulator/repository/sql"
	"SQLSimulator/utils"
)

func main() {
	logger := logging.Logger()
	//conf := config.GetConfig(logger)
	conf := make(map[string]string)
	entityStruct := entity.SQLSimulatorInterfaces{
		Input:      api.NewApiData(),
		Repository: repsql.NewSQLData(),
	}

	entityStruct.Initiate(conf, logger)
	for {
		entityStruct.ProcceedSave(logger)
		entityStruct.ProcceedUpdate(logger)
		utils.RandSleep()

	}
}
