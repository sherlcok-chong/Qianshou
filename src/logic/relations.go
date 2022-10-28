package logic

import (
	"Qianshou/src/dao"
	db "Qianshou/src/dao/postgres/sqlc"
	"Qianshou/src/global"
	"Qianshou/src/model/reply"
	"Qianshou/src/model/request"
	"Qianshou/src/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type relations struct {
}

func (relations) GetUserRelations(c *gin.Context, userID int64) ([]reply.Relations, errcode.Err) {
	result, err := dao.Group.DB.GetAllRelations(c, userID)
	if err != nil {
		global.Logger.Error(err.Error())
		return []reply.Relations{}, errcode.ErrServer
	}
	rly := make([]reply.Relations, 0)
	for _, v := range result {
		t := reply.Relations{
			UserID: v.Tid,
			State:  string(v.RelationType),
			Type:   "relationship",
		}
		rly = append(rly, t)
	}
	return rly, nil
}

func (relations) UpdateRelations(c *gin.Context, updateRelations request.UpdateRelations) (reply.Relations, errcode.Err) {
	result, err := dao.Group.DB.UpdateRelation(c, &db.UpdateRelationTypeParams{
		RelationType: db.Relationtype(updateRelations.State),
		Fid:          updateRelations.FID,
		Tid:          updateRelations.TID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.Relations{}, errcode.ErrServer
	}
	return reply.Relations{
		UserID: result.Tid,
		State:  string(result.RelationType),
		Type:   "relationship",
	}, nil
}
