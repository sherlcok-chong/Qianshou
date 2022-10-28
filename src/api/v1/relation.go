package v1

import (
	"Qianshou/src/logic"
	"Qianshou/src/model/request"
	"Qianshou/src/pkg/app"
	"Qianshou/src/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type relation struct {
}

func (relation) GetMyRelation(c *gin.Context) {
	rly := app.NewResponse(c)
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	r, mErr := logic.Group.Relations.GetUserRelations(c, userID)
	rly.ReplyList(mErr, int64(len(r)), r)
}

func (relation) UpdateRelation(c *gin.Context) {
	rly := app.NewResponse(c)
	fID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	tID, err := strconv.ParseInt(c.Param("other_user_id"), 10, 64)
	if err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	params := &request.UpdateRelations{}
	if err := c.BindJSON(params); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	params.FID = fID
	params.TID = tID
	r, mErr := logic.Group.Relations.UpdateRelations(c, *params)
	rly.Reply(mErr, r)
}
