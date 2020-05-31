// author: ashing
// time: 2020/5/31 3:47 下午
// mail: axingfly@gmail.com
// Less is more.

package db

import (
	"fmt"

	"github.com/ronething/pan/db/mysql"
)

//UserSignUp 用户注册 DB
func UserSignUp(username string, passwd string) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_user " +
		"(`user_name`, `user_pwd`) values (?,?)")
	if err != nil {
		fmt.Printf("Failed to insert, err:%s\n", err.Error())
		return false
	}

	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Printf("Failed to insert, err: %s\n", err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}

	return false

}

func UserSignIn(username string, encpwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"select * from tbl_user " +
			"wher user_name=? limit 1")
	if err != nil {
		fmt.Printf("Failed to select, err:%s\n", err.Error())
		return false
	}

	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Printf("Failed to select, err: %s\n", err.Error())
		return false
	} else if rows == nil {
		fmt.Printf("username not found:%s\n", username)
	}

	pRows := mysql.ParseRows(rows)

	return false

}
