package http_srv

import (
	"github.com/giskook/lorawan/base"
	"net/http"
	"context"
	"time"
	"strings"
)

type user_del struct{ 
	Code int `json:"code"`
	Desc string `json:"desc"`
}

func (h *HttpSrv) handler_web_user_del(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	r.ParseForm()
	defer r.Body.Close()
	id := r.Form.Get(WEB_USER_DEL_PARA_ID)
	if id == ""{
		EncodeErrResponse(w, base.ERROR_HTTP_LACK_PARAMTERS)
		return
	}
	id = strings.Replace(id, "'","''", -1)

	ctx := r.Context()
	var cancel context.CancelFunc
	ctx ,cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	c_result := make(chan *base.DBResult)
	go h.db.UserDel(ctx, id, c_result)
	for{
	select{
	case <- ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
	case result:= <- c_result:
		if result.Err != nil{ 
			EncodeErrResponse(w, result.Err.(*base.LorawanError))
			return
		}else{
			EncodeErrResponse(w, base.ERROR_NONE)
		}
	}
}
}