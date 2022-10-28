package setting

import (
	"Qianshou/src/dao"
	"Qianshou/src/dao/postgres"
	"Qianshou/src/global"
	"fmt"
)

type mDao struct {
}

// Init 持久化层初始化
func (m mDao) Init() {
	fmt.Println(global.PvSettings.Postgresql.SourceName)
	dao.Group.DB = postgres.Init(global.PvSettings.Postgresql.SourceName)
}
