package xerrors

import "fmt"

//（1）错误码小于100(不包含15,40,41错误码)的调用错误，这种错误一般是由于用户的请求不符合各种基本校验而引起的。用户遇到这些错误的返回首先检查应用的权限、频率等情况，然后参照文档检验一下传入的参数是否完整且合法。
//（1）错误码为40,41的错误；40主要是必填参数没有传入报错，41主要是传入的参数格式不对报错。:2.
//（2）错误码大于100或者等于15且子错误码（sub_code）是"isv."开头【( 错误码 > 100 or 错误码 = 15 )  and (isv开头)】的调用错误:

const (
	ErrCodeUploadFail                  = 3  //上传失败
	ErrCodeAppCallLimited              = 4  //app 调用次数限制
	ErrCodeServiceCurrentlyUnavailable = 10 //服务暂不可用
	ErrCodeInsufficientUserPermissions = 12 //用户权限不足
	ErrCodeInsufficientAppPermissions  = 13 //app权限不足
	ErrCodeRemoteServiceError          = 15 //服务错误
	ErrCodeMissingMethod               = 21 //
	ErrCodeInvalidMethod               = 22 //非法api method
	ErrCodeMissingRequiredArguments    = 40 //缺少必要参数
	ErrCodeInvalidArguments            = 41 //参数格式不对
	ErrCodeMissingAccessToken          = 42 //缺少access_token
	ErrCodeInvalidAccessToken          = 43 //非法access_token
	ErrCodeInternalServerError         = 500
)

//isv.###-not-exist:*** 根据***查询不到###
//isv.missing-parameter:*** 缺少必要的参数***
//isv.invalid-parameter:***  参数***无效，格式不对、非法值、越界等
//isv.invalid-permission 权限不够、非法访问
//isv.parameters-mismatch:***-and-### 传入的参数***和###不匹配，两者有一定的对应关系
//isv.***-service-error:### 调用***服务返回false，业务逻辑错误，###表示具体的错误信息

//isp.***-service-unavailable	调用后端服务***抛异常，服务不可用	ISP	是
//isp.remote-service-error	连接远程服务错误	ISP	是
//isp.remote-service-timeout	连接远程服务超时	ISP	是
//isp.remote-connection-error	远程连接错误	ISP	是
//isp.null-pointer-exception	空指针异常错误	ISP	否
//isp.top-parse-error	api解析错误（出现了未被明确控制的异常信息）	ISP	否
//isp.top-remote-connection-timeout	top平台连接后端服务超时	ISP	是
//isp.top-remote-connection-error	top平台连接后端服务错误，找不到服务	ISP	是
//isp.top-mapping-parse-error	top-mapping转换出错，主要是由于传入参数格式不对	ISP	否
//isp.unknown-error	top平台连接后端服务抛未知异常信息	ISP	是

type GateError struct {
	err     error
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	SubMsg  string `json:"sub_msg"`
	SubCode string `json:"sub_code"`
}

//msg,sub_code,sub_msg
func NewGateError(code int, args ...string) *GateError {

	var ret = &GateError{
		Code: code,
	}
	length := len(args)

	switch {
	case length >= 1:
		ret.Msg = args[0]
	case length >= 2:
		ret.SubCode = args[1]
	case length >= 3:
		ret.SubMsg = args[2]
	}

	return ret
}

func (data *GateError) SetError(err error) *GateError {
	data.err = err
	return data
}

func (data *GateError) Error() string {

	return fmt.Sprintf("code:%d \n msg: %s \n sub_code: %s \n sub_msg: %s err:%s", data.Code, data.Msg, data.SubCode, data.SubMsg, data.err)
}
