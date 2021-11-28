package web

import (
	"encoding/json"
	"net/http"
	"microblog/microblog/domain/queryservice/dto"
	"microblog/microblog/usecase"

	"github.com/gorilla/sessions"
)

const COOKIE_SESSION_ID = "session-id"

var store = sessions.NewCookieStore([]byte("AF2YYDS3YXxpGTh0qUTJgH5BAUEa7zv1"))

func RegisterAccountHandler(interactor usecase.RegisterAccount) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.RegisterAccountInputData

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := interactor(&input, func(output *usecase.RegisterAccountOutputData) {
			session, err := store.New(r, COOKIE_SESSION_ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			session.Values["user_id"] = int64(output.UserID)

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})

		if err != nil {
			handleError(w, err)
		}
	}
}

func LoginHandler(interactor usecase.Login) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.LoginInputData

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := interactor(&input, func(output *usecase.LoginOutputData) {
			session, err := store.New(r, COOKIE_SESSION_ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			session.Values["user_id"] = int64(output.UserID)

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})

		if err != nil {
			handleError(w, err)
		}
	}
}

type PostTweetInput struct {
	Message string `json:"message"`
}

func PostTweetHandler(interactor usecase.PostTweet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input PostTweetInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		currentUserID, ok := r.Context().Value("user_id").(int64)
		if !ok {
			http.Error(w, "invalid context", http.StatusInternalServerError)
			return
		}

		err := interactor(&usecase.PostTweetInputData{
			CurrentUserID: currentUserID,
			Message:       input.Message,
		}, func() {})
		if err != nil {
			handleError(w, err)
		}
	}
}

type FollowInput struct {
	UserID int64 `json:"user_id"`
}

func FollowHandler(interactor usecase.Follow) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input FollowInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		currentUserID, ok := r.Context().Value("user_id").(int64)
		if !ok {
			http.Error(w, "invalid context", http.StatusInternalServerError)
			return
		}

		err := interactor(&usecase.FollowInputData{
			FollowerUserId: currentUserID,
			FolloweeUserId: input.UserID,
		}, func() {})
		if err != nil {
			handleError(w, err)
		}
	}
}

type GetTweetListResponse struct {
	Tweets []*dto.Tweet `json:"tweets"`
}

func GetTweetListHandler(interactor usecase.GetTweetList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.GetTweetListInput

		userID, ok := r.Context().Value("user_id").(int64)
		if !ok {
			http.Error(w, "invalid context", http.StatusInternalServerError)
			return
		}

		input.CurrentUserID = userID

		err := interactor(&input, func(tweets []*dto.Tweet) {
			res, err := json.Marshal(GetTweetListResponse{tweets})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		})
		if err != nil {
			handleError(w, err)
		}
	}
}

type GetUserListResponse struct {
	Users []*dto.User `json:"users"`
}

func GetUserListHandler(interactor usecase.GetUserList) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.GetUserListInput

		userID, ok := r.Context().Value("user_id").(int64)
		if !ok {
			http.Error(w, "invalid context", http.StatusInternalServerError)
			return
		}

		input.CurrentUserID = userID

		err := interactor(&input, func(users []*dto.User) {
			res, err := json.Marshal(GetUserListResponse{users})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		})
		if err != nil {
			handleError(w, err)
		}

	}
}
