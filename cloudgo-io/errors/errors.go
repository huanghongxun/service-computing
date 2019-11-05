package errors

import (
	. "errors"
)

var (
	codes = make(map[error]ErrorCode)

	ErrNotFound                = New("资源不存在")
	ErrMethodNotAllow          = New("方法不被允许")
	ErrBadRequest              = New("请求发生错误")
	ErrInvalidRequestParameter = New("无效的请求参数")
	ErrTooManyRequests         = New("请求过于频繁")
	ErrNotImplemented          = New("接口未实现")
	ErrUnknownQuery            = New("未知的查询类型")
	ErrInvalidParent           = New("无效的父级节点")
	ErrNotAllowDeleteWithChild = New("含有子级，不能删除")
	ErrResourceExists          = New("资源已经存在")
	ErrResourceNotAllowDelete  = New("资源不允许删除")

	ErrNoPerm         = New("无访问权限")
	ErrNoResourcePerm = New("无资源的访问权限")

	ErrInvalidUserName = New("无效的用户名")
	ErrInvalidPassword = New("无效的密码")
	ErrInvalidUser     = New("无效的用户")
	ErrUserDisable     = New("用户被禁用")
	ErrUserNotEmptyPwd = New("密码不允许为空")

	// login
	ErrLoginNotAllowModifyPwd = New("不允许修改密码")
	ErrLoginInvalidOldPwd     = New("旧密码不正确")
	ErrLoginInvalidVerifyCode = New("无效的验证码")

	ErrInternal = New("内部错误")
)

func init() {
	newBadRequestError(ErrBadRequest)
	newBadRequestError(ErrInvalidRequestParameter)
	newErrorCode(ErrNotFound, 404, ErrNotFound.Error(), 404)
	newErrorCode(ErrMethodNotAllow, 405, ErrMethodNotAllow.Error(), 405)
	newErrorCode(ErrTooManyRequests, 429, ErrTooManyRequests.Error(), 429)
	newErrorCode(ErrNotImplemented, 501, ErrNotImplemented.Error(), 501)
	newBadRequestError(ErrUnknownQuery)
	newBadRequestError(ErrInvalidParent)
	newBadRequestError(ErrNotAllowDeleteWithChild)
	newBadRequestError(ErrResourceExists)
	newBadRequestError(ErrResourceNotAllowDelete)

	newErrorCode(ErrNoPerm, 9999, ErrNoPerm.Error(), 401)
	newErrorCode(ErrNoResourcePerm, 401, ErrNoResourcePerm.Error(), 401)

	newBadRequestError(ErrInvalidUserName)
	newBadRequestError(ErrInvalidPassword)
	newBadRequestError(ErrInvalidUser)
	newBadRequestError(ErrUserDisable)
	newBadRequestError(ErrUserNotEmptyPwd)

	newBadRequestError(ErrLoginNotAllowModifyPwd)
	newBadRequestError(ErrLoginInvalidOldPwd)
	newBadRequestError(ErrLoginInvalidVerifyCode)

	newErrorCode(ErrInternal, 500, ErrInternal.Error(), 500)
}

// ErrorCode 错误码
type ErrorCode struct {
	Code           int
	Message        string
	HTTPStatusCode int
}

// newErrorCode 设定错误码
func newErrorCode(err error, code int, message string, status ...int) {
	errCode := ErrorCode{
		Code:    code,
		Message: message,
	}
	if len(status) > 0 {
		errCode.HTTPStatusCode = status[0]
	}
	codes[err] = errCode
}

// FromErrorCode 获取错误码
func FromErrorCode(err error) (ErrorCode, bool) {
	v, ok := codes[err]
	return v, ok
}

// newBadRequestError 创建请求错误
func newBadRequestError(err error) {
	newErrorCode(err, 400, err.Error(), 400)
}

// newUnauthorizedError 创建未授权错误
func newUnauthorizedError(err error) {
	newErrorCode(err, 401, err.Error(), 401)
}

// newInternalServerError 创建服务器错误
func newInternalServerError(err error) {
	newErrorCode(err, 500, err.Error(), 500)
}
