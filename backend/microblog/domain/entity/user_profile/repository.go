package userprofile

type UserProfileRepository interface {
	Create(userAccount *UserProfile) (*UserProfile, error)
}
