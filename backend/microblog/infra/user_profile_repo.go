package infra

import (
	userprofile "microblog/microblog/domain/entity/user_profile"

	"github.com/jmoiron/sqlx"
)

type UserProfileRepo struct {
	db *sqlx.DB
}

func NewUserProfiletRepo(db *sqlx.DB) *UserProfileRepo {
	return &UserProfileRepo{db}
}

func (r *UserProfileRepo) Create(profile *userprofile.UserProfile) (*userprofile.UserProfile, error) {
	result, err := r.db.NamedExec("insert into user_profiles (user_id, display_name, description) values (:user_id, :display_name, :description)", profile)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdProfile := userprofile.UserProfile{}
	err = r.db.Get(&createdProfile, "select * from user_profiles where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &createdProfile, nil
}
