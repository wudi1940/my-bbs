package processer

import (
	"fmt"
	"my-bbs/cmd/client"
)

var (
	dbClient = client.DbClient
)

func CreateUser(user client.User) (int64, error) {

	cc := dbClient.DBClient

	//准备sql语句
	stmt, err := cc.Prepare("INSERT INTO user_info (`user_id`, `nickname`, `email`, `phone`, `password`, `avatar`, `ctime`, `mtime`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.UserId, user.Nickname, user.Email, user.Phone, user.Password, user.Avatar, user.Ctime, user.Mtime)
	if err != nil {
		fmt.Println("Exec fail")
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}
