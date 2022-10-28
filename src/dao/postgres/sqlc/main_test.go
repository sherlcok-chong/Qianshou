package db_test

import (
	"Qianshou/src/setting"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setting.Group.Config.Init()
	setting.Group.Dao.Init()
	os.Exit(m.Run())
}
