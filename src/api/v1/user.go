package v1

import (
	"Qianshou/src/global"
	"Qianshou/src/logic"
	"Qianshou/src/model/request"
	"Qianshou/src/pkg/app"
	"Qianshou/src/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) GetAllUser(c *gin.Context) {
	rly := app.NewResponse(c)
	result, err := logic.Group.User.GetAllUsers(c)
	rly.ReplyList(err, int64(len(result)), result)
}

func (user) CreateUser(c *gin.Context) {
	rly := app.NewResponse(c)
	params := &request.User{}
	if err := c.BindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	if len(params.Name) > global.PbSettings.Rule.UsernameLenMax || len(params.Name) < global.PbSettings.Rule.UsernameLenMin {
		rly.Reply(errcode.ErrParamsNotValid)
		return
	}
	result, err := logic.Group.User.CreateUser(c, params.Name)
	rly.Reply(err, result)
}
