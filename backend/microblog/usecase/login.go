package usecase

type LoginInputData struct {
	Username string
	Password string
}

type LoginOutputData struct {
	UserID int64
}

type Login func(input *LoginInputData, presenter func(output *LoginOutputData)) error
