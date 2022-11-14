package entity

type SQLSimulatorInterfaces struct {
	Input      Input
	Repository Repository
}

func NewSQLSimulatorInterfaces() *SQLSimulatorInterfaces {
	return &SQLSimulatorInterfaces{}
}

type Entity struct {
	Word string
	Sql  struct {
		Id           int
		Data         string //TODO check not bigger than 40 char
		isSuccesfull bool
	}
}

func NewEntity() *Entity {
	return &Entity{}
}
