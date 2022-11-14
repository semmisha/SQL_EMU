package repsql

const (
	SaveSingleData   = "Insert into test_cdc_table (test_data) values ('$1')"
	UpdateSingleData = "Update test_cdc_table set (test_data = $1) where test_id = $2;"
	DeleteSingleData = "Delete from test_cdc_table where test_id = $1; "
	GetAllID         = " Select test_id from test_cdc_table"
)
