package logic

type group struct {
	User      user
	Relations relations
}

var Group = new(group)
