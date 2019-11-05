package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/huanghongxun/cloudgo-io/errors"
	"github.com/huanghongxun/cloudgo-io/schema"
	"net/http"
)

func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.ErrInvalidRequestParameter
	}
	return nil
}

func ResSuccessText(c *gin.Context, v string) {
	ResText(c, http.StatusOK, v)
}

func ResText(c *gin.Context, status int, v string) {
	c.Data(status, "text/html; charset=utf-8", []byte(v))
	c.Abort()
}

func ResSuccessJSON(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

func ResError(c *gin.Context, err error, status ...int) {
	statusCode := 500
	errItem := schema.HTTPErrorItem{
		Code:    500,
		Message: "服务器发生错误",
	}

	if errCode, ok := errors.FromErrorCode(err); ok {
		errItem.Code = errCode.Code
		errItem.Message = errCode.Message
		statusCode = errCode.HTTPStatusCode
	}

	if len(status) > 0 {
		statusCode = status[0]
	}

	ResJSON(c, statusCode, schema.HTTPError{Error: errItem})
}
