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

		logger.Fatal(err, err1, username, password, url, db)
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
	fmt.Println(entity.Sql.Id, entity.Sql.Data, entity.Word)
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

	return true

}

func (S *SQLData) DeleteSingleData(entity *entity.Entity, logger *logrus.Logger) (isSuccesfull bool) {
	//TODO implement me
	panic("implement me")
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
