package dto

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	DisplayName string `db:"display_name" json:"display_name"`
	Description string `json:"description"`
	Following   bool   `json:"following"`
	FollowedBy  bool   `db:"followed_by" json:"followed_by"`
}
