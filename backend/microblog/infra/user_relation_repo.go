package infra

import (
	userrelation "microblog/microblog/domain/entity/user_relation"

	"github.com/jmoiron/sqlx"
)

type UserRelationRepo struct {
	db *sqlx.DB
}

func NewUserRelationRepo(db *sqlx.DB) *UserRelationRepo {
	return &UserRelationRepo{db}
}

func (r *UserRelationRepo) Add(relation *userrelation.UserRelation) error {
	_, err := r.db.NamedExec("insert ignore into user_relations (follower_user_id, followee_user_id) values (:follower_user_id, :followee_user_id)", relation)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRelationRepo) Remove(relation *userrelation.UserRelation) error {
	_, err := r.db.NamedExec(
		"delete from user_relations where follower_user_id = :follower_user_id and followee_user_id = :followee_user_id",
		relation,
	)
	if err != nil {
		return err
	}

	return nil
}
