package router

import (
	v1 "Qianshou/src/api/v1"
	mid "Qianshou/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinLogger())
	root := r.Group("users")
	{
		root.GET("", v1.Group.User.GetAllUser)
		root.POST("", v1.Group.User.CreateUser)
		root.GET("/:user_id/relationships", v1.Group.Relation.GetMyRelation)
		root.PUT("/:user_id/relationships/:other_user_id", v1.Group.Relation.UpdateRelation)
	}
	return r
}
