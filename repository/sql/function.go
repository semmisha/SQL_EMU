package repsql

import (
	"SQLSimulator/entity"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sirupsen/logrus"
)

func (S *SQLData) Initiate(config map[string]string, logger *logrus.Logger) {
	var (
		url = config["url"]
		//port      = config["port"]
		db       = config["db"]
		username = config["username"]
		password = config["password"]
	)

	connString := fmt.Sprintf("sqlserver://%v:%v@%v?database=%v", username, password, url, db)

	pool, err := sql.Open("sqlserver", connString)
	err1 := pool.Ping()
	if err != nil || err1 != nil {

		logger.Fatal(err, err1)
	}
	S.Connection = pool

}

func (S *SQLData) UpdateSingleData(entity *entity.Entity, logger *logrus.Logger) (isSuccesfull bool) {
	var (
		conn = S.Connection
	)

	_, err := conn.Exec("Update test_cdc_table set test_data = @Data where test_id = @ID;", sql.Named("Data", entity.Word), sql.Named("ID", entity.Sql.Id))
	if err != nil {
		logger.Error(err)
		return false
	}
	logger.Infof("Entry updated, id: %v", entity.Sql.Id)
	return true

}

func (S *SQLData) SaveSingleData(entity *entity.Entity, logger *logrus.Logger) (isSuccesfull bool) {

	var (
		conn = S.Connection
	)

	_, err := conn.Exec("Insert into test_cdc_table (test_data) values (@Data)", sql.Named("Data", entity.Word))
	if err != nil {
		logger.Error(err)
		return false
	}
	logger.Infof("Entry inserted, id: %v , data: ", entity.Sql.Id, entity.Word)
	return true

}

func (S *SQLData) DeleteSingleData(entity *entity.Entity, logger *logrus.Logger) (isSuccesfull bool) {
	var (
		conn = S.Connection
	)

	_, err := conn.Exec("Delete from test_cdc_table where test_id = @ID; ", sql.Named("ID", entity.Sql.Id))
	if err != nil {
		logger.Error(err)
		return false
	}
	logger.Infof("Entry deleted, id: %v", entity.Sql.Id)
	return true
}

func (S *SQLData) GetAllID(logger *logrus.Logger) (AllID []int, isSuccesfull bool) {

	var (
		conn = S.Connection
	)

	rows, err := conn.Query(GetAllID)
	if err != nil {
		logger.Error(err)
		return nil, false
	}
	for rows.Next() {
		var id int
		err1 := rows.Scan(&id)
		if err1 != nil {
			logger.Error(err1)
		}
		AllID = append(AllID, id)

	}
	return AllID, true

}
