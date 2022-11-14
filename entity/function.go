package entity

import (
	"SQLSimulator/utils"
	"github.com/sirupsen/logrus"
)

func (S SQLSimulatorInterfaces) Initiate(config map[string]string, logger *logrus.Logger) {
	S.Input.Initiate(config, logger)
	S.Repository.Initiate(config, logger)
}

func (S *SQLSimulatorInterfaces) ProcceedSave(logger *logrus.Logger) (isSuccesfull bool) {

	var (
		input      = S.Input
		repository = S.Repository
		entity     = NewEntity()
	)

	if word, ok := input.GetSingleData(logger); ok {
		entity.Word = word
		if ok = repository.SaveSingleData(entity, logger); ok {
			return true
		}
	}
	return false

	// Read config from env
	// Connect to DB
	// get random word from api
	// insert to DB
	// sleep for rand time

	// Read config from env
	// Connect to DB
	// get random word from api
	// get all id from db save to slice
	// rand len slice
	// update to DB
	// sleep for rand time

	// Read config from env
	// Connect to DB
	// get all id from db save to slice
	// rand len slice
	// delete from DB
	// sleep for rand time

}

func (S *SQLSimulatorInterfaces) ProcceedUpdate(logger *logrus.Logger) (isSuccesfull bool) {

	var (
		input      = S.Input
		repository = S.Repository
	)
	entity := NewEntity()
	if IDs, ok := repository.GetAllID(logger); ok {
		entity.Sql.Id = utils.RandSlice(IDs)
	}

	if word, ok := input.GetSingleData(logger); ok {
		entity.Word = word
		if ok = repository.UpdateSingleData(entity, logger); ok {
			return true
		}
	}
	return false

	// Read config from env
	// Connect to DB
	// get random word from api
	// insert to DB
	// sleep for rand time

	// Read config from env
	// Connect to DB
	// get random word from api
	// get all id from db save to slice
	// rand len slice
	// update to DB
	// sleep for rand time

	// Read config from env
	// Connect to DB
	// get all id from db save to slice
	// rand len slice
	// delete from DB
	// sleep for rand time

}
