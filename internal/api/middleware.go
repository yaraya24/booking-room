package api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/yaraya24/book-meeting-room/internal/domain"
	"github.com/yaraya24/book-meeting-room/internal/pkg/logging"

	"github.com/sirupsen/logrus"
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

// LoggingMiddlware logs the method, path, request-id, response status and
// duration of every request that passes through it.
type LoggingMiddlware struct{}

func NewLoggingMiddleware() LoggingMiddlware {
	return LoggingMiddlware{}
}

func (l LoggingMiddlware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := generateRequestID()
		w.Header().Set("X-Request-Id", requestID)

		ctx := logging.WithFields(r.Context(), logging.Fields{"request_id": requestID})
		r = r.WithContext(ctx)

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		logging.FromContext(ctx).WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status":      recorder.status,
			"duration_ms": time.Since(start).Milliseconds(),
		}).Info("handled request")
	})
}

// statusRecorder wraps http.ResponseWriter to capture the status code written
// by downstream handlers so it can be logged.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// generateRequestID returns a random 16 character hex string used to
// correlate log lines for a single request.
func generateRequestID() string {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "unknown"
	}
	return hex.EncodeToString(buf)
}
