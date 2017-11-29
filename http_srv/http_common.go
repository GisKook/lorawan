package http_srv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type ctx_key int

type Code int

const (
	HTTP_REP_SUCCESS        Code    = 0
	HTTP_REP_LACK_PARAMETER Code    = 1
	HTTP_REP_INTERAL_ERROR  Code    = 2
	xkey                    ctx_key = 0
)

var HTTP_REQUEST_DESC []string = []string{
	"成功",
	"缺少参数",
	"服务器内部错误"}

func (c Code) Desc() string {
	return HTTP_REQUEST_DESC[c]
}

type GeneralResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func EncodeGeneralResponse(code Code) string {
	general_response := &GeneralResponse{
		Code: int(code),
		Desc: code.Desc(),
	}

	resp, _ := json.Marshal(general_response)

	return string(resp)
}

func EncodeErrResponse(code int, desc string) string {
	gr := &GeneralResponse{
		Code: code,
		Desc: desc,
	}

	resp, _ := json.Marshal(gr)

	return string(resp)
}

func RecordReq(r *http.Request) {
	v, e := httputil.DumpRequest(r, true)
	if e != nil {
		log.Println(e.Error())
		return
	}
	log.Println(string(v))
}

// MarshalJson 把对象以json格式放到response中
func marshal_json(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func unmarshal_json(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	return nil
}
