// File:    error_msg
// Version: 1.0.0
// Creator: JoeLang
// Date:    2020/8/31 7:42
// DESC:

package error_msg

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000... 用户表错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_EXPIRED      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_FORMAT_WRONG = 1007
)

var codeMsg = map[int]string{
	200: "Success",
	500: "Fail",

	1001: "User Already Exists.",
	1002: "Password Wrong.",
	1003: "User Does Not Exist.",
	1004: "Token Does Not Exist",
	1005: "Token Already Expired.",
	1006: "Token Wrong.",
	1007: "Token Format Wrong.",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
