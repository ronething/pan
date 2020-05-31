// author: ashing
// time: 2020/5/30 10:31 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"net/http"

	"github.com/ronething/pan/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)         // 上传
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)  // 上传成功提示
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)      // 获取元数据
	http.HandleFunc("/file/query", handler.FileQueryHandler)       // 批量获取元数据
	http.HandleFunc("/file/download", handler.DownloadHandler)     // 下载
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler) // 重命名
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)     // 删除

	http.HandleFunc("/user/signup", handler.SignupHandler) // 用户注册
	err := http.ListenAndServe("127.0.0.1:8887", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err: %s", err.Error())
	}
}
