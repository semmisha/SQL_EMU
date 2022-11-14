package entity

import "github.com/sirupsen/logrus"

type Input interface {
	Initialization
	GetSingleData
}

type Repository interface {
	Initialization
	UpdateSingleData(entity *Entity, logger *logrus.Logger) (isSuccesfull bool)
	SaveSingleData(entity *Entity, logger *logrus.Logger) (isSuccesfull bool)
	DeleteSingleData(entity *Entity, logger *logrus.Logger) (isSuccesfull bool)
	GetAllID(logger *logrus.Logger) (AllID []int, isSuccesfull bool)
}

type GetSingleData interface {
	GetSingleData(logger *logrus.Logger) (word string, isSuccesfull bool)
}

type Initialization interface {
	Initiate(config map[string]string, logger *logrus.Logger)
}
