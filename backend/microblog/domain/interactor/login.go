package interactor

import (
	"fmt"
	"microblog/microblog/domain"
	useraccount "microblog/microblog/domain/entity/user_account"
	"microblog/microblog/usecase"
)

func Login(uar useraccount.UserAccountRepository) usecase.Login {
	return func(input *usecase.LoginInputData, presenter func(*usecase.LoginOutputData)) error {
		ua, err := uar.FindByUsername(input.Username)
		switch err {
		case nil:
		case domain.ErrNotFound:
			return fmt.Errorf("%w: username or password is wrong.", domain.ErrAuthentication)
		default:
			return err
		}

		if ok := ua.VerifyPassword(input.Password); !ok {
			return fmt.Errorf("%w: username or password is wrong.", domain.ErrAuthentication)
		}

		presenter(&usecase.LoginOutputData{int64(ua.ID)})

		return nil
	}
}
