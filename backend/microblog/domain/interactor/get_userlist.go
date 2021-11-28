package interactor

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/domain/queryservice"
	"microblog/microblog/domain/queryservice/dto"
	"microblog/microblog/usecase"
)

func GetUserList(uqs queryservice.UserQueryService) usecase.GetUserList {
	return func(input *usecase.GetUserListInput, presenter func(users []*dto.User)) error {
		userID := useraccount.UserID(input.CurrentUserID)

		users, err := uqs.FindAll(userID)
		if err != nil {
			return err
		}

		presenter(users)

		return nil
	}
}
