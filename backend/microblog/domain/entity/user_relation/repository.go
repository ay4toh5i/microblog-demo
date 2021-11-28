package userrelation

type UserRelationRepository interface {
	Add(*UserRelation) error
	Remove(*UserRelation) error
}
