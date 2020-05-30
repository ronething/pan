// author: ashing
// time: 2020/5/30 11:11 上午
// mail: axingfly@gmail.com
// Less is more.

package meta

import "sort"

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

//UpdateFileMeta 新增/更新文件元数据
func UpdateFileMeta(f FileMeta) { // TODO: 考虑线程安全问题
	fileMetas[f.FileSha1] = f
}

//GetFileMeta 根据 hash 获取文件元数据
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
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

//RemoveFileMeta 删除元数据
func RemoveFileMeta(filehash string) {
	delete(fileMetas, filehash) // TODO: 考虑线程安全问题
}
