package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

const ( //静态变量
	userName   = "root"
	pwd        = "123456"
	ipCode     = "127.0.0.1"
	sqlPort    = "3306"
	dbName     = "ds"
	driverName = "mysql"
)

func initialDB() {
	db, err = sql.Open(driverName, userName+":"+pwd+"@tcp("+ipCode+":"+sqlPort+")/"+dbName+"?charset=utf8") //打开一个数据库的连接，直接通过db调用相关方法就可以执行sql语句
	if err != nil {
		fmt.Println(err)
	}
	println("连接成功")
}
