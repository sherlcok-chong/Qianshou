package setting

type group struct {
	Dao    mDao
	Log    log
	Config config
}

var Group = new(group)

func AllInit() {
	Group.Config.Init()
	Group.Dao.Init()
	Group.Log.Init()
}
