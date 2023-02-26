package api

import (
	"github.com/HelliWrold1/quaver/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

/*
* @Author:
* @Date:
* @Description:
 */

type Response struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_message"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}
