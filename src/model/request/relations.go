package request

type UpdateRelations struct {
	FID   int64  `json:"fid"`
	TID   int64  `json:"tid"`
	State string `json:"state" binding:"required"`
}
