// author: ashing
// time: 2020/5/30 5:15 下午
// mail: axingfly@gmail.com
// Less is more.

package db

import (
	"database/sql"
	"fmt"

	"github.com/ronething/pan/db/mysql"
)

//OnFileUploadFinished 文件上传完成,保存 meta 元数据
func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into `tbl_file` " +
			"(`file_sha1`,`file_name`, `file_size`, `file_addr`,`status`) " +
			"values (?,?,?,?,1)",
	)
	if err != nil {
		fmt.Printf("Failed to prepare statement, err:%s\n", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}

	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMeta 从数据库获取文件元数据
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mysql.DBConn().Prepare(
		"select file_sha1, file_addr,file_name,file_size " +
			"from tbl_file where file_sha1=? and status=1 limit 1",
	)

	if err != nil {
		fmt.Printf("Failed to prepare statement, err:%s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(
		&tfile.FileHash, &tfile.FileAddr,
		&tfile.FileName, &tfile.FileSize,
	)
	if err != nil {
		fmt.Printf("scan err: %s\n", err.Error())
		return nil, err
	}
	return &tfile, nil

}
