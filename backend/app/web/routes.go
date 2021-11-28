package web

import (
	"microblog/microblog/domain/interactor"
	"microblog/microblog/infra"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func routes(r *mux.Router, db *sqlx.DB) {

	uar := infra.NewUserAccountRepo(db)
	upr := infra.NewUserProfiletRepo(db)
	ur := infra.NewUserRelationRepo(db)
	tr := infra.NewTweetRepo(db)

	tqs := infra.NewTweetQS(db)
	uqs := infra.NewUserQS(db)

	r.HandleFunc("/users", auth(GetUserListHandler(interactor.GetUserList(uqs)))).Methods("GET")
	r.HandleFunc("/users", RegisterAccountHandler(interactor.RegisterAccount(uar, upr, ur))).Methods("POST")
	r.HandleFunc("/followings", auth(FollowHandler(interactor.Follow(uar, ur)))).Methods("POST")
	r.HandleFunc("/tweets", auth(PostTweetHandler(interactor.PostTweet(tr)))).Methods("POST")
	r.HandleFunc("/tweets", auth(GetTweetListHandler(interactor.GetTweetList(tqs)))).Methods("GET")
	r.HandleFunc("/login", LoginHandler(interactor.Login(uar))).Methods("POST")
}
