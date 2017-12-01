package http_srv

import (
	"github.com/giskook/lorawan/base"
	"net/http"
)

func (h *HttpSrv) handler_web_user_main(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	EncodeErrResponse(w, base.ERROR_NONE)
}
