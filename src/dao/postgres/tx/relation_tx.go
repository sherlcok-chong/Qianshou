package tx

import (
	db "Qianshou/src/dao/postgres/sqlc"
	"context"

	"github.com/jackc/pgx/v4"
)

func (store *SqlStore) UpdateRelation(c context.Context, arg *db.UpdateRelationTypeParams) (db.Relationship, error) {
	var rly db.Relationship
	err := store.execTx(c, func(queries *db.Queries) error {
		_, err := queries.GetRelationType(c, &db.GetRelationTypeParams{
			Fid: arg.Fid,
			Tid: arg.Tid,
		})
		if err != nil && err != pgx.ErrNoRows {
			return err
		}
		t, err2 := queries.GetRelationType(c, &db.GetRelationTypeParams{
			Fid: arg.Tid,
			Tid: arg.Fid,
		})
		if err2 != nil && err2 != pgx.ErrNoRows {
			return err
		}
		var r *db.Relationship
		var e error
		if arg.RelationType == "like" { //如果喜欢操作
			if err2 == pgx.ErrNoRows || t.RelationType != "like" { // 如果对方现在不喜欢自己
				if err == pgx.ErrNoRows { //如果没有记录
					r, e = queries.CreateRelations(c, &db.CreateRelationsParams{
						Fid:          arg.Fid,
						Tid:          arg.Tid,
						RelationType: "like",
					})
					if e != nil {
						return e
					}
					rly = db.Relationship{
						Fid:          r.Fid,
						Tid:          r.Tid,
						RelationType: r.RelationType,
					}
				} else { //如果有记录
					e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
						RelationType: "like",
						Fid:          arg.Fid,
						Tid:          arg.Tid,
					})
					if e != nil {
						return e
					}
					// TODO
					rly = db.Relationship{
						Fid:          arg.Fid,
						Tid:          arg.Tid,
						RelationType: "like",
					}
				}
			} else { //如果对方现在喜欢自己
				if err == pgx.ErrNoRows { //如果没有记录
					r, e = queries.CreateRelations(c, &db.CreateRelationsParams{
						Fid:          arg.Fid,
						Tid:          arg.Tid,
						RelationType: "matched",
					})
					if e != nil {
						return e
					}
					e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
						RelationType: "matched",
						Fid:          arg.Tid,
						Tid:          arg.Fid,
					})
					if e != nil {
						return e
					}
					rly = db.Relationship{
						Fid:          r.Fid,
						Tid:          r.Tid,
						RelationType: r.RelationType,
					}
				} else {
					e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
						RelationType: "matched",
						Fid:          arg.Fid,
						Tid:          arg.Tid,
					})
					if e != nil {
						return e
					}
					e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
						RelationType: "matched",
						Fid:          arg.Tid,
						Tid:          arg.Fid,
					})
					if e != nil {
						return e
					}
					rly = db.Relationship{
						Fid:          arg.Fid,
						Tid:          arg.Tid,
						RelationType: "matched",
					}
				}
			}
		} else { // 如果不喜欢操作
			if err == pgx.ErrNoRows { // 如果没有记录
				r, e = queries.CreateRelations(c, &db.CreateRelationsParams{
					Fid:          arg.Fid,
					Tid:          arg.Tid,
					RelationType: "dislike",
				})
				if e != nil {
					return e
				}
				rly = db.Relationship{
					Fid:          r.Fid,
					Tid:          r.Tid,
					RelationType: r.RelationType,
				}
			} else { // 如果有记录
				e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
					RelationType: "dislike",
					Fid:          arg.Fid,
					Tid:          arg.Tid,
				})
				if e != nil {
					return e
				}
				if t.RelationType == "matched" {
					e = queries.UpdateRelationType(c, &db.UpdateRelationTypeParams{
						RelationType: "like",
						Fid:          arg.Tid,
						Tid:          arg.Fid,
					})
					if e != nil {
						return e
					}
				}
				rly = db.Relationship{
					Fid:          arg.Fid,
					Tid:          arg.Tid,
					RelationType: "dislike",
				}
			}

		}

		return nil
	})
	return rly, err
}
