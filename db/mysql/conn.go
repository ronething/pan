// author: ashing
// time: 2020/5/30 5:11 下午
// mail: axingfly@gmail.com
// Less is more.

package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000) // 同时连接数
	err := db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, err: %s\n", err.Error())
		os.Exit(1)
	}

}

//DBConn 返回数据库对象
func DBConn() *sql.DB {
	return db

}
