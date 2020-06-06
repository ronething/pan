// author: ashing
// time: 2020/5/31 3:47 下午
// mail: axingfly@gmail.com
// Less is more.

package db

import (
	"fmt"

	"github.com/ronething/pan/db/mysql"
)

type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

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

//UserSignIn 用户登录 判断密码是否一致
func UserSignIn(username string, encpwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"select * from tbl_user " +
			"where user_name=? limit 1")
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

	pRows := mysql.ParseRows(rows) // TODO: review
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}

	fmt.Printf("verify username and pwd error")
	return false

}

//GetUserInfo 查询用户信息
func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mysql.DBConn().Prepare(
		"select user_name, signup_at from tbl_user where user_name=? limit 1",
	)
	if err != nil {
		fmt.Printf("prepare sql err:%s\n", err.Error())
		return user, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}

	return user, nil

}

//UpdateToken 更新用户登录 token
func UpdateToken(username string, token string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`, `user_token`) values (?,?)",
	)
	if err != nil {
		fmt.Printf("prepare sql err:%s\n", err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Printf("exec err: %s\n", err.Error())
		return false
	}

	return true
}
