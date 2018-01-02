package http_srv

import (
	"time"
	"context"
	"github.com/giskook/lorawan/base"
	"net/http"
	"html/template"
	"strings"
	"log"
)

const(
	USER_MANAGEMENT string = "um"
	DEVICE_MANAGEMENT string = "dm"
)

type device_management struct{
	Bind string
}
type sys_management struct{ 
	UserManagement string
}

type mine struct{
	Setting string
}

type header struct{
	Name string 
	DM *device_management
	SYS *sys_management
	Mine *mine
}

type main struct{
	Title string
	Header *header
}

func (h *HttpSrv) gen_menu(name, auth string) * header{
	head := new(header)
	head.Name = name
	if strings.Contains(auth, DEVICE_MANAGEMENT) {
		head.DM = new(device_management)
		head.DM.Bind=DEVICE_MANAGEMENT
	}
	if strings.Contains(auth, USER_MANAGEMENT) {
		head.SYS = new(sys_management)
		head.SYS.UserManagement =USER_MANAGEMENT
	}

	head.Mine = new(mine)
	head.Mine.Setting = "mine"

	return head
}

func (h *HttpSrv) handler_web_user_main(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
			EncodeErrResponse(w,base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	RecordReq(r)
	r.ParseForm()
	defer r.Body.Close()

	ctx := r.Context()
	token, ok := ctx.Value(auth_key).(JwtAuth)
	if !ok {
		http.NotFound(w, r)
		return
	}

	var cancel context.CancelFunc
	ctx ,cancel = context.WithTimeout(ctx, time.Duration(h.conf.Http.TimeOut)*time.Second)
	defer cancel()

	user_info := make(chan *base.DBResult)
	go h.db.UserInfoGet(token.Userid, user_info)
	select{
	case <-ctx.Done():
		EncodeErrResponse(w, base.ERROR_HTTP_TIMEOUT)
	case info := <-user_info: 
	    if info.Err != nil{ 
			EncodeErrResponse(w,base.NewErr(info.Err, base.ERR_DB_USER_INFO_GET_CODE, base.ERR_DB_USER_INFO_GET_DESC))
		}else{
		m := main{
			Title:info.Extra.(string),
			Header:h.gen_menu(info.Extra.(string), token.Auth),
		}
		tmpl := template.Must(template.ParseFiles("./web/tmpl/main.tmpl", "./web/tmpl/common/header.tmpl", "./web/tmpl/user/user_dlg_add.tmpl", "./web/tmpl/user/user_dlg_modify.tmpl", "./web/tmpl/device/device_dlg_add.tmpl"))
		err := tmpl.Execute(w ,&m)
		
		panic(err)
		}
	}
}
