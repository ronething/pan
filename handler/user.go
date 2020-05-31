// author: ashing
// time: 2020/5/31 3:51 下午
// mail: axingfly@gmail.com
// Less is more.

package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/ronething/pan/db"
	"github.com/ronething/pan/util"
)

const (
	pwd_salt = "*#890"
)

//SignupHandler 用户注册
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	if len(username) < 4 || len(passwd) < 6 {
		w.Write([]byte("Invalid params"))
		return
	}

	encode_pwd := util.Sha1([]byte(passwd + pwd_salt))

	if suc := db.UserSignUp(username, encode_pwd); suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}

	return

}

//SignInHandler 用户登录
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	// 1、校验用户名及密码

	// 2、生成访问凭证

	// 3、登录成功重定向到首页

}
