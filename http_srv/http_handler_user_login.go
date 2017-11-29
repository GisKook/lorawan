package http_srv

import (
	"fmt"
	"github.com/giskook/lorawan/base"
	"log"
	"net/http"
	"time"
)

func (h *HttpSrv) handler_web_user_login(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_INTERAL_ERROR))
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
		fmt.Fprint(w, EncodeGeneralResponse(HTTP_REP_LACK_PARAMETER))
		return
	}

	auth, id, err := h.db.UserValid(user, passwd)
	if err != nil {
		log.Println((err.(*base.LorawanError)).Desc())
		return
	}

	err = h.set_cookie(w, r, auth, id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/web/user/main", 307)
}

func (h *HttpSrv) handler_web_user_logout(w http.ResponseWriter, r *http.Request) {
	RecordReq(r)
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	return
}
