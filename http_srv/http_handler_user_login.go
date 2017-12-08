package http_srv

import (
	"github.com/giskook/lorawan/base"
	"log"
	"net/http"
	"time"
)

func (h *HttpSrv) handler_web_user_login(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	user := r.Form.Get(WEB_USER_LOGIN_PARA_USER)
	passwd := r.Form.Get(WEB_USER_LOGIN_PARA_PASSWD)
	//var login Login
	if user == "" ||
		passwd == "" {
			EncodeErrResponse(w, base.ERROR_HTTP_LACK_PARAMTERS)
		return
	}

	auth, id, err := h.db.UserValid(user, passwd)
	if err != nil {
		log.Println((err.(*base.LorawanError)).Desc())
		EncodeErrResponse(w, err.(*base.LorawanError))
		return
	}

	err = h.set_cookie(w, r, auth, id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	EncodeErrResponse(w, base.ERROR_NONE)
}

func (h *HttpSrv) handler_web_user_logout(w http.ResponseWriter, r *http.Request) {
	RecordReq(r)
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	http.Redirect(w, r, "/", 307)
	return
}
