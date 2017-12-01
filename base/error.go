package base

import ()

type LorawanError struct {
	Err      error
	Code     int
	Describe string
}

func (e *LorawanError) Error() string {
	return e.Err.Error()
}

func (e *LorawanError) Desc() string {
	return e.Describe
}

func NewErr(err error, code int, desc string) *LorawanError {
	return &LorawanError{
		Err:      err,
		Code:     code,
		Describe: desc,
	}
}

const (
	ERR_NONE_CODE int = 0
	ERR_NONE_DESC string = "成功"

	ERR_HTTP_LACK_PARAMTERS_CODE int = 1 
	ERR_HTTP_LACK_PARAMTERS_DESC string = "缺少参数"

	ERR_HTTP_INNER_PANIC_CODE int = 2 
	ERR_HTTP_INNER_PANIC_DESC string = "内部错误"

	ERR_DB_BEGIN_TRANSCATION_CODE int    = 100
	ERR_DB_BEGIN_TRANSCATION_DESC string = "[DB]开启事务失败"

	ERR_DB_COMMIT_TRANSCATION_CODE int    = 101
	ERR_DB_COMMIT_TRANSCATION_DESC string = "[DB]提交事务失败"

	ERR_DB_QUERY_FAIL_CODE int    = 102
	ERR_DB_QUERY_FAIL_DESC string = "[DB]查询失败"

	ERR_USER_NOT_FOUND_CODE      int    = 200
	ERR_USER_NOT_FOUND_DESC      string = "用户不存在"
	ERR_USER_UNVALID_PASSWD_CODE int    = 201
	ERR_USER_UNVALID_PASSWD_DESC string = "密码错误"
)

var ( 
	ERROR_HTTP_LACK_PARAMTERS *LorawanError = NewErr(nil, ERR_HTTP_LACK_PARAMTERS_CODE, ERR_HTTP_LACK_PARAMTERS_DESC)
	ERROR_HTTP_INNER_PANIC *LorawanError = NewErr(nil, ERR_HTTP_INNER_PANIC_CODE, ERR_HTTP_INNER_PANIC_DESC)
	ERROR_NONE *LorawanError = NewErr(nil, ERR_NONE_CODE, ERR_NONE_DESC)
)