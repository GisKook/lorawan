package http_srv

import (
	"github.com/giskook/lorawan/base"
	"net/http"
	"strconv"
	"context"
	"time"
)

type user_search struct{ 
	Code int `json:"code"`
	Desc string `json:"desc"`
	Users []*base.User `json:"users"`
	UsersCount int `json:"users_count"`
	PageIndex int `json:"page_index"`
	PageSize int `json:"page_size"`
	collect int
}

func (h *HttpSrv) handler_web_user_search(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()
	id := r.Form.Get(WEB_USER_MODIFY_PARA_ID)
	alias := r.Form.Get(WEB_USER_MODIFY_PARA_ALIAS)
	page_index,_ := strconv.Atoi(r.Form.Get(WEB_USER_MODIFY_PARA_PAGE_INDEX))
	page_size,_ := strconv.Atoi(r.Form.Get(WEB_USER_MODIFY_PARA_PAGE_SIZE))

	ctx := r.Context()
	var cancel context.CancelFunc
	ctx ,cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	c_users := make(chan *base.DBResult)
	c_users_count := make(chan *base.DBResult)
	go h.db.UserSearch(ctx, id, alias,page_size, page_size*page_index, c_users)
	go h.db.UserSearchCount(ctx, id, alias, c_users_count)

	result := user_search{Code:base.ERR_NONE_CODE, Desc:base.ERR_NONE_DESC}
	result.PageIndex = page_index
	result.PageSize = page_size
	for{
	select{
	case <- ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
	case users := <- c_users:
		if users.Err != nil{ 
			EncodeErrResponse(w, users.Err.(*base.LorawanError))
			return
		}
		result.Users = users.Extra.([]*base.User)
		result.collect++
		if result.collect == 2{
			marshal_json(w, result)
			return 
		}
	case users_count := <-c_users_count:
		if users_count.Err != nil{
			EncodeErrResponse(w, users_count.Err.(*base.LorawanError))
			return
		}
		result.UsersCount = users_count.Extra.(int)
		result.collect++
		if result.collect == 2{
			marshal_json(w, result)
			return
		}
	}
	}
}