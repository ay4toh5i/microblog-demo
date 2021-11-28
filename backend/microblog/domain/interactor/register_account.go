package interactor

import (
	"errors"
	"fmt"
	"microblog/microblog/domain"
	useraccount "microblog/microblog/domain/entity/user_account"
	userprofile "microblog/microblog/domain/entity/user_profile"
	userrelation "microblog/microblog/domain/entity/user_relation"
	"microblog/microblog/usecase"
)

func RegisterAccount(
	uar useraccount.UserAccountRepository,
	upr userprofile.UserProfileRepository,
	ur userrelation.UserRelationRepository,
) usecase.RegisterAccount {
	return func(input *usecase.RegisterAccountInputData, presenter func(output *usecase.RegisterAccountOutputData)) error {
		account, err := useraccount.NewUserAccount(input.Username, input.Password)
		if err != nil {
			return err
		}

		_, err = uar.FindByUsername(input.Username)
		if err == nil {
			return fmt.Errorf("%w: the username is already registered.", domain.ErrInvalidRequest)
		}
		if ok := errors.Is(err, domain.ErrNotFound); !ok {
			return err
		}

		account, err = uar.Create(account)
		if err != nil {
			return err
		}

		profile, err := userprofile.NewUserProfile(int64(account.ID), input.DisplayName, input.Description)

		profile, err = upr.Create(profile)
		if err != nil {
			return err
		}

		err = ur.Add(&userrelation.UserRelation{
			FollowerUserID: account.ID,
			FolloweeUserId: account.ID,
		})
		if err != nil {
			return err
		}

		presenter(&usecase.RegisterAccountOutputData{
			UserID: int64(account.ID),
		})

		return nil
	}
}
