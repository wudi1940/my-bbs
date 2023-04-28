package processer

import (
	"my-bbs/cmd/client"
	"testing"
)

func TestCreateUser(t *testing.T) {
	client.DbClient.InitDBClient()

	user := client.User{
		UserId:   0,
		Nickname: "TestGod",
		Email:    "",
		Phone:    "",
		Password: "password",
		Avatar:   "",
		Ctime:    "",
		Mtime:    "",
	}

	CreateUser(user)

}
