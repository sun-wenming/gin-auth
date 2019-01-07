package e

const (
	SUCCESS                      = 200
	ERROR                        = 500
	INVALID_PARAMS_WITHOUT_TOKEN = 401

	// 认证
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	// --- 客户端错误
	ERROR_USER_GET_INFO        = 40001
	ERROR_USER_REG_NAME        = 40002
	ERROR_USER_NAME_EXIST      = 40003
	ERROR_PHONE_NOT_VALID      = 40004
	ERROR_PHONE_CODE_EXPIRED   = 40005
	ERROR_PHONE_CODE_NOT_VALID = 40006
	ERROR_USER_NAME_NOT_EXIST  = 40007
	ERROR_USER_INFO_EMPTY      = 40008
	ERROR_INVALID_PARAMS       = 40009
	// --- end
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_INVALID_PARAMS:           "请求参数错误",
	INVALID_PARAMS_WITHOUT_TOKEN:   "不存在Token参数",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_USER_GET_INFO:            "40001: 获取到用户信息失败.",
	ERROR_USER_REG_NAME:            "40002: 用户输入格式错误.",
	ERROR_USER_NAME_EXIST:          "40003: 用户名已存在.",
	ERROR_PHONE_NOT_VALID:          "40004: 手机号验证失败.",
	ERROR_PHONE_CODE_EXPIRED:       "40005: 验证码已过期.",
	ERROR_PHONE_CODE_NOT_VALID:     "40006: 验证码验证失败.",
	ERROR_USER_NAME_NOT_EXIST:      "40007: 用户名不存在.",
	ERROR_USER_INFO_EMPTY:          "40008: 账号无对应用户信息.",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
