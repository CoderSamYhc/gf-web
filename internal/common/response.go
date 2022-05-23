package common

import (
	"context"
	"fmt"
	"gf-web/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)



var CodeMap = g.MapIntStr{
	consts.ERROR : "系统错误：%v",
}

type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type List struct {
	List []*interface{} `json:"list"`
	Page int `json:"page"`
	Size int `json:"size"`
	Total int `json:"total"`
}

type Response struct {
	Ctx context.Context
}

func (r *Response) Success(data interface{}) {
	_ = g.RequestFromCtx(r.Ctx).Response.WriteJson(Result{
		Code: http.StatusOK,
		Message: "ok",
		Data: data,
	})
}

func (r *Response) Error(code int, err string) {
	var (
		msg = fmt.Sprintf(CodeMap[code], err)
	)
	_ = g.RequestFromCtx(r.Ctx).Response.WriteJson(Result{
		Code: code,
		Message: msg,
	})
}
