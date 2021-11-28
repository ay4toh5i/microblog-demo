package useraccount

type UserAccountRepository interface {
	Find(userID UserID) (*UserAccount, error)
	FindByUsername(username string) (*UserAccount, error)
	Create(userAccount *UserAccount) (*UserAccount, error)
}
