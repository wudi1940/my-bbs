package client

import "testing"

func TestDbClient_InitDBClient(t *testing.T) {
	db := DbClient{}

	db.InitDBClient()
}
