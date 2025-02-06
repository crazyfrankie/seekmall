package response

import (
	"github.com/crazyfrankie/seekmall/app/pkg/constants"
	"net/http"

	"github.com/crazyfrankie/gem/gerrors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: constants.Success.BizStatusCode(),
		Msg:  constants.Success.BizMessage(),
		Data: data,
	})
}

func Error(c *gin.Context, err error) {
	if bizErr, ok := gerrors.FromBizStatusError(err); ok {
		c.JSON(http.StatusOK, Response{
			Code: bizErr.BizStatusCode(),
			Msg:  bizErr.BizMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: constants.InternalServer.BizStatusCode(),
		Msg:  constants.InternalServer.BizMessage(),
	})
}
