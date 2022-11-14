package main

import (
	"SQLSimulator/config"
	"SQLSimulator/controller/api"
	"SQLSimulator/entity"
	"SQLSimulator/logging"
	"SQLSimulator/repository/sql"
	"SQLSimulator/utils"
	"strconv"
	"sync"
)

func main() {
	logger := logging.Logger()
	conf := config.GetConfig(logger)
	sleepDuration, err := strconv.Atoi(conf["sleepDuration"])
	if err != nil {
		logger.Fatal("Cant ATOI sleep duration", err)
	}
	//entityStruct := entity.SQLSimulatorInterfaces{
	//	Input:      api.NewApiData(),
	//	Repository: repsql.NewSQLData(),
	//}
	//entityStruct.Initiate(conf, logger)
	var vg sync.WaitGroup
	vg.Add(1)
	go func() {

		entityStruct := entity.SQLSimulatorInterfaces{
			Input:      api.NewApiData(),
			Repository: repsql.NewSQLData(),
		}
		entityStruct.Initiate(conf, logger)
		for {
			entityStruct.ProcceedSave(logger)
			utils.RandSleep(sleepDuration)
		}

	}()
	vg.Add(1)
	go func() {

		entityStruct := entity.SQLSimulatorInterfaces{
			Input:      api.NewApiData(),
			Repository: repsql.NewSQLData(),
		}
		entityStruct.Initiate(conf, logger)
		for {
			entityStruct.ProcceedUpdate(logger)
			utils.RandSleep(sleepDuration)
		}

	}()
	vg.Add(1)
	go func() {

		entityStruct := entity.SQLSimulatorInterfaces{
			Input:      api.NewApiData(),
			Repository: repsql.NewSQLData(),
		}
		entityStruct.Initiate(conf, logger)
		for {
			entityStruct.ProcceedDelete(logger)
			utils.RandSleep(sleepDuration + 300)
		}

	}()
	vg.Wait()

}
