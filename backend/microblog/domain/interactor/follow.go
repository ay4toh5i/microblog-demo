package interactor

import (
	useraccount "microblog/microblog/domain/entity/user_account"
	userrelation "microblog/microblog/domain/entity/user_relation"
	"microblog/microblog/usecase"
)

func Follow(uar useraccount.UserAccountRepository, urr userrelation.UserRelationRepository) usecase.Follow {
	return func(input *usecase.FollowInputData, presenter func()) error {
		followerUserID := useraccount.UserID(input.FollowerUserId)
		followeeUserID := useraccount.UserID(input.FolloweeUserId)

		_, err := uar.Find(followeeUserID)
		if err != nil {
			return err
		}

		err = urr.Add(&userrelation.UserRelation{
			FollowerUserID: followerUserID,
			FolloweeUserId: followeeUserID,
		})
		if err != nil {
			return err
		}

		presenter()

		return nil
	}
}
