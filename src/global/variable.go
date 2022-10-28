package global

import (
	"Qianshou/src/model/config"

	"github.com/0RAJA/Rutils/pkg/logger"
)

var (
	Logger     *logger.Log    // 日志
	PbSettings config.Public  // Public配置
	PvSettings config.Private // Private配置
)
