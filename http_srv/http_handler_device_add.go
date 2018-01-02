package http_srv

import (
	"github.com/giskook/lorawan/base"
	"net/http"
	"context"
	"time"
)

func (h *HttpSrv) handler_web_device_add(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	r.ParseForm()
	defer r.Body.Close()

	id := r.Form.Get(WEB_DEVICE_ADD_PARA_ID)
	if id == ""{
		EncodeErrResponse(w, base.ERROR_HTTP_LACK_PARAMTERS)
		return
	}
	
	ctx := r.Context()
	var cancel context.CancelFunc
	ctx ,cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	device_add_result := make(chan *base.DBResult)
	go h.db.DeviceAdd(ctx, id, device_add_result)
	select{
	case <-ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
	case result := <-device_add_result:
		if result.Err != nil{
			EncodeErrResponse(w, result.Err.(*base.LorawanError))
			return
		}
		EncodeErrResponse(w, base.ERROR_NONE)
	}
}