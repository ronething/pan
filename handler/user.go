// author: ashing
// time: 2020/5/31 3:51 下午
// mail: axingfly@gmail.com
// Less is more.

package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}

	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPassword := util.Sha1([]byte(password + pwd_salt))

	// 1、校验用户名及密码
	pwdChecked := db.UserSignIn(username, encPassword)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}
	// 2、生成访问凭证
	token := GenToken(username)
	upRes := db.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}

	// 3、登录成功重定向到首页
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			UserName string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			UserName: username,
			Token:    token,
		},
	}

	w.Write(resp.JSONBytes())

}

// GenToken : 生成token 40 位
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
