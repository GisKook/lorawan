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
	ERR_COMMON_NOT_CAPTURE_CODE int = 999
	ERR_COMMON_NOT_CAPTURE_DESC string = "未捕获的错误"

	ERR_NONE_CODE int = 0
	ERR_NONE_DESC string = "成功"

	ERR_HTTP_LACK_PARAMTERS_CODE int = 1 
	ERR_HTTP_LACK_PARAMTERS_DESC string = "缺少参数"

	ERR_HTTP_INNER_PANIC_CODE int = 2 
	ERR_HTTP_INNER_PANIC_DESC string = "内部错误"

	ERR_HTTP_TIMEOUT_CODE int = 3
	ERR_HTTP_TIMEOUT_DESC string = "超时"

	ERR_DB_BEGIN_TRANSCATION_CODE int    = 100
	ERR_DB_BEGIN_TRANSCATION_DESC string = "[DB]开启事务失败"

	ERR_DB_COMMIT_TRANSCATION_CODE int    = 101
	ERR_DB_COMMIT_TRANSCATION_DESC string = "[DB]提交事务失败"

	ERR_DB_QUERY_FAIL_CODE int    = 102
	ERR_DB_QUERY_FAIL_DESC string = "[DB]查询失败"
	ERR_DB_USER_INFO_GET_CODE int = 103
	ERR_DB_USER_INFO_GET_DESC string = "用户信息查询失败"


	ERR_USER_NOT_FOUND_CODE      int    = 200
	ERR_USER_NOT_FOUND_DESC      string = "用户不存在"
	ERR_USER_UNVALID_PASSWD_CODE int    = 201
	ERR_USER_UNVALID_PASSWD_DESC string = "密码错误"

	ERR_USER_ALREADY_EXIST_CODE int = 202
	ERR_USER_ALREADY_EXIST_DESC string = "用户已存在"

	ERR_USER_ROLE_ADD_CODE int = 203
	ERR_USER_ROLE_ADD_DESC string = "添加角色失败"

	ERR_USER_QUERY_CODE int = 204
	ERR_USER_QUERY_DESC string = "查询用户失败"

	ERR_USER_DEL_CODE int = 205
	ERR_USER_DEL_DESC string = "删除用户失败"

	ERR_USER_DEL_NO_ID_CODE int = 206
	ERR_USER_DEL_NO_ID_DESC string = "没有指定要删除的用户"
)

var ( 
	ERROR_HTTP_LACK_PARAMTERS *LorawanError = NewErr(nil, ERR_HTTP_LACK_PARAMTERS_CODE, ERR_HTTP_LACK_PARAMTERS_DESC)
	ERROR_HTTP_INNER_PANIC *LorawanError = NewErr(nil, ERR_HTTP_INNER_PANIC_CODE, ERR_HTTP_INNER_PANIC_DESC)
	ERROR_HTTP_TIMEOUT *LorawanError = NewErr(nil, ERR_HTTP_TIMEOUT_CODE, ERR_HTTP_TIMEOUT_DESC)
	ERROR_NONE *LorawanError = NewErr(nil, ERR_NONE_CODE, ERR_NONE_DESC)
	ERROR_NOT_CAPTURE *LorawanError = NewErr(nil, ERR_COMMON_NOT_CAPTURE_CODE, ERR_COMMON_NOT_CAPTURE_DESC)
)