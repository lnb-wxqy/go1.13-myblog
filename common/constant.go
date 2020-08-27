package common

const (

	//api
	// 用户相关
	USER_GROUP = "/api/auth"
	REGISTER   = "/register"
	LOGIN      = "/login"
	INFO       = "/info"

	//文章相关
	CATEGORY_GROUP = "/categories"
	CREATE         = "/create"
	UPDATE         = "/update"
	SHOW           = "/show"
	DELETE         = "/delete"

	//business status code
	USER_IS_NOT_EXIST            = "1001"
	USER_IS_EXIST                = "1002"
	USERNAME_IS_NULL             = "1003"
	PASSWORD_IS_NOT_CORRECT      = "1004"
	PASSWORD_LENGTH_LESS_SIX     = "1005"
	TELEPHONE_FORMAT_ERROR       = "1006"
	GENERATE_FROM_PASSWORD       = "1007" //加密错误
	REGISTER_FAILED              = "1008"
	STATUS_INTERNAL_SERVER_ERROR = "1009" //系统错误
	AUTHORIZATION_ERROR          = "1010"

	TOKEN_IS_INVALID = "-1001"
)

var StatusText = map[string]string{
	USER_IS_NOT_EXIST:            "用户不存在",
	USER_IS_EXIST:                "用户已存在",
	USERNAME_IS_NULL:             "用户名不能为空",
	PASSWORD_IS_NOT_CORRECT:      "密码错误",
	PASSWORD_LENGTH_LESS_SIX:     "密码长度不能少于6位",
	TELEPHONE_FORMAT_ERROR:       "手机号格式不正确",
	GENERATE_FROM_PASSWORD:       "密码加密错误",
	TOKEN_IS_INVALID:             "token无效",
	STATUS_INTERNAL_SERVER_ERROR: "系统错误",
	AUTHORIZATION_ERROR:          "Authorization 错误",
}
