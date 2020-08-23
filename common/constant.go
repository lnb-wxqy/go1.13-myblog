package common

const (

	//api
	GROUP    = "/api/auth"
	REGISTER = "/register"
	LOGIN    = "/login"
	INFO     = "/info"

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
	USER_IS_NOT_EXIST:            "user is not exist",
	USER_IS_EXIST:                "user is already exist",
	USERNAME_IS_NULL:             "username is null",
	PASSWORD_IS_NOT_CORRECT:      "password is not correct",
	PASSWORD_LENGTH_LESS_SIX:     "password length is less 6",
	TELEPHONE_FORMAT_ERROR:       "telephone format error",
	GENERATE_FROM_PASSWORD:       "generate from password",
	TOKEN_IS_INVALID:             "token is invalid",
	STATUS_INTERNAL_SERVER_ERROR: "internal server Error",
	AUTHORIZATION_ERROR:          "Authorization error",
}
