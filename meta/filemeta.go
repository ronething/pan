// author: ashing
// time: 2020/5/30 11:11 上午
// mail: axingfly@gmail.com
// Less is more.

package meta

import (
	"sort"

	"github.com/ronething/pan/db"
)

//FileMeta 文件元数据模型
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// key 存储文件 hash
var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMetaDB 新增/更新文件元数据到数据库
func UpdateFileMetaDB(f *FileMeta) bool { // TODO: 考虑线程安全问题
	return db.OnFileUploadFinished(f.FileSha1, f.FileName, f.FileSize, f.Location)
}

//GetFileMetaDB 从数据库获取文件元数据
func GetFileMetaDB(fileSha1 string) (*FileMeta, error) {
	tfile, err := db.GetFileMeta(fileSha1)
	if err != nil {
		return nil, err
	}
	tMeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return &tMeta, nil
}

//GetLastFileMetas 获取数据
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	// ByUploadTime []FileMeta 根据 time 时间排序
	sort.Sort(ByUploadTime(fMetaArray))

	return fMetaArray[0:count]
}

//GetLastFileMetasDB 批量从数据库获取元数据
func GetLastFileMetasDB(limit int) ([]FileMeta, error) {
	tfiles, err := db.GetFileMetaList(limit)
	if err != nil {
		return make([]FileMeta, 0), err
	}
	tfilesMap := make([]FileMeta, len(tfiles))
	for i := 0; i < len(tfilesMap); i++ {
		tfilesMap[i] = FileMeta{
			FileSha1: tfiles[i].FileHash,
			FileName: tfiles[i].FileName.String,
			FileSize: tfiles[i].FileSize.Int64,
			Location: tfiles[i].FileAddr.String,
		}
	}

	return tfilesMap, nil

}

//RemoveFileMeta 删除元数据
func RemoveFileMeta(filehash string) {
	delete(fileMetas, filehash) // TODO: 考虑线程安全问题
}
