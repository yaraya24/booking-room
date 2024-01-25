package api

import (
	"context"
	"net/http"

	"github.com/yaraya24/book-meeting-room/internal/domain"
	"github.com/yaraya24/book-meeting-room/internal/pkg/logging"
)

type AuthenticationMiddlware struct {
	Repo AuthenticationRepo
}

type AuthenticationRepo interface {
	FindUser(ctx context.Context, username string) (domain.User, error)
}

func NewAuthenticationMiddleware(repo AuthenticationRepo) AuthenticationMiddlware {
	return AuthenticationMiddlware{Repo: repo}
}

func (a AuthenticationMiddlware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := a.Repo.FindUser(r.Context(), username)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if user.Password != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", user.ID)
		ctx = logging.WithFields(ctx, logging.Fields{"user_id": user.ID})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
