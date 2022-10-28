package dao

import "Qianshou/src/dao/postgres"

type group struct {
	DB postgres.DB
}

var Group = new(group)
