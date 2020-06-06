// author: ashing
// time: 2020/6/6 3:18 下午
// mail: axingfly@gmail.com
// Less is more.

package handler

import "net/http"

//AuthorizedInterceptor 鉴权 middleware
func AuthorizedInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username := r.Form.Get("username")
		token := r.Form.Get("token")

		// 验证 token 有效性
		if len(username) < 3 || IsTokenValid(token) {
			http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
			return
		}

		// 验证通过
		h(w, r)
	}

}

//IsTokenValid token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断 token 的时效性，是否过期
	// TODO: 从数据库表 tbl_user_token 查询 username 对应的token信息
	// TODO: 对比两个 token 是否一致
	return true
}
