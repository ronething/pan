// author: ashing
// time: 2020/6/5 3:06 下午
// mail: axingfly@gmail.com
// Less is more.

package db

import (
	"fmt"
	"time"

	"github.com/ronething/pan/db/mysql"
)

//UserFile 用户文件返回结构体
type UserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

func (u UserFile) Scan(src interface{}) error {
	panic("implement me")
}

//OnUserFileUploadFinished 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`, `file_sha1`, " +
			"`file_name`,`file_size`, `upload_at`) values (?,?,?,?,?)",
	)
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		return false
	}

	return true

}

//QueryUserFileMetas 查询用户文件列表
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mysql.DBConn().Prepare(
		"select file_sha1, file_name, file_size, upload_at, last_update " +
			"from tbl_user_file where  user_name=? limit ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil {
		return nil, err
	}

	var userFiles []UserFile

	for rows.Next() {
		ufile := UserFile{}
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSize,
			&ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			fmt.Printf("err is:%s\n", err.Error())
			break
		}
		userFiles = append(userFiles, ufile)
	}

	return userFiles, nil
}
