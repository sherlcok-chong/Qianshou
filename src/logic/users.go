package logic

import (
	"Qianshou/src/dao"
	db "Qianshou/src/dao/postgres/sqlc"
	"Qianshou/src/global"
	"Qianshou/src/model/reply"
	"Qianshou/src/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) CreateUser(c *gin.Context, username string) (reply.User, errcode.Err) {
	mUser, err := dao.Group.DB.CreateUsers(c, &db.CreateUsersParams{
		Name:     username,
		UserType: "user",
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return reply.User{}, errcode.ErrServer
	}
	return reply.User{
		ID:   mUser.ID,
		Name: mUser.Name,
		Type: string(mUser.UserType),
	}, nil
}

func (user) GetAllUsers(c *gin.Context) ([]reply.User, errcode.Err) {
	users, err := dao.Group.DB.GetAllUser(c)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ErrServer
	}
	rly := make([]reply.User, 0)
	for _, v := range users {
		t := reply.User{
			ID:   v.ID,
			Name: v.Name,
			Type: string(v.UserType),
		}
		rly = append(rly, t)
	}
	return rly, nil
}
