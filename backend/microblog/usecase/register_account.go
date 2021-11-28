package usecase

type RegisterAccountInputData struct {
	Username    string
	Password    string
	DisplayName string `json:"display_name"`
	Description string
}

type RegisterAccountOutputData struct {
	UserID int64
}

type RegisterAccount func(input *RegisterAccountInputData, presenter func(output *RegisterAccountOutputData)) error
