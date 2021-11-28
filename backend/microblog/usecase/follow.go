package usecase

type FollowInputData struct {
	FollowerUserId int64
	FolloweeUserId int64
}

type Follow func(input *FollowInputData, presenter func()) error
