// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreateRelations(ctx context.Context, arg *CreateRelationsParams) (*Relationship, error)
	CreateUsers(ctx context.Context, arg *CreateUsersParams) (*User, error)
	GetAllRelations(ctx context.Context, fid int64) ([]*Relationship, error)
	GetAllUser(ctx context.Context) ([]*User, error)
	GetRelationType(ctx context.Context, arg *GetRelationTypeParams) (*Relationship, error)
	UpdateRelationType(ctx context.Context, arg *UpdateRelationTypeParams) error
}

var _ Querier = (*Queries)(nil)
