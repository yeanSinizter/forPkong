package db_test

import (
	"eiei/clean_architecture/db"
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, teardown := db.InitSql()

	defer teardown()
}
