package consts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 封装常用的常量，以及响应封装数据
const (
	NilMessage          = ""            // 空字符串
	TokenExpireDuration = time.Hour * 2 // Token 的过期时间
	CtxAdminIDKey       = "ctxAdminIdKey"
	CtxTokenKey         = "ctxTokenKey"
)

// RespCode 响应码封装
type RespCode int32

const (
	CodeSuccess RespCode = 1000 + iota
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var codeMsg = map[RespCode]string{
	CodeSuccess:      "成功",
	CodeServerBusy:   "服务繁忙",
	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

func (c RespCode) GetMsg() string {
	msg, ok := codeMsg[c]
	if !ok {
		msg = codeMsg[c]
	}
	return msg
}

// ResponseData 返回的响应信息结构体
type ResponseData struct {
	Code    RespCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseError 错误响应
func ResponseError(ctx *gin.Context, code RespCode) {
	var response = &ResponseData{
		Code:    code,
		Message: code.GetMsg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}

// ResponseSuccess 成功响应
func ResponseSuccess(ctx *gin.Context, message string, data interface{}) {
	var response = &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.GetMsg(),
		Data:    data,
	}
	if message != "" {
		response.Message = message
	}
	ctx.JSON(http.StatusOK, response)
}
