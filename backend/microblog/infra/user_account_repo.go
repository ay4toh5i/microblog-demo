package infra

import (
	"database/sql"
	"microblog/microblog/domain"
	useraccount "microblog/microblog/domain/entity/user_account"

	"github.com/jmoiron/sqlx"
)

type UserAccountRepo struct {
	db *sqlx.DB
}

func NewUserAccountRepo(db *sqlx.DB) *UserAccountRepo {
	return &UserAccountRepo{db}
}

func (r *UserAccountRepo) Find(userID useraccount.UserID) (*useraccount.UserAccount, error) {
	account := useraccount.UserAccount{}
	err := r.db.Get(&account, "select * from user_accounts where id = ?", userID)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}

	return &account, nil
}

func (r *UserAccountRepo) FindByUsername(username string) (*useraccount.UserAccount, error) {
	account := useraccount.UserAccount{}
	err := r.db.Get(&account, "select * from user_accounts where username = ?", username)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}

	return &account, nil
}

func (r *UserAccountRepo) Create(userAccount *useraccount.UserAccount) (*useraccount.UserAccount, error) {
	result, err := r.db.NamedExec("insert into user_accounts (username, password) values (:username, :password)", userAccount)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createdAccount := useraccount.UserAccount{}
	err = r.db.Get(&createdAccount, "select * from user_accounts where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &createdAccount, nil
}
