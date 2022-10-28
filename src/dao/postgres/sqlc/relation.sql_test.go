package db_test

import (
	"Qianshou/src/dao"
	db "Qianshou/src/dao/postgres/sqlc"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueries_GetRelationType(t *testing.T) {
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "good",
			f: func() {
				arg := &db.GetRelationTypeParams{
					Fid: 1,
					Tid: 2,
				}
				r, err := dao.Group.DB.GetRelationType(context.Background(), arg)
				require.NoError(t, err)
				require.NotNil(t, r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f()
		})
	}
}
