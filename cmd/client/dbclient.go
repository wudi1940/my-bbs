package client

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DbClient = &dbClientStruct{}

type dbClientStruct struct {
	DBClient *sql.DB
}

func (d *dbClientStruct) InitDBClient() {
	db, _ := sql.Open("mysql", "")
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		panic("open database fail")
	}
	fmt.Println("connnect success")
	d.DBClient = db
}
