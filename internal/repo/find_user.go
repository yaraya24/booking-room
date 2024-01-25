package repo

import (
	"context"
	"database/sql"

	"github.com/yaraya24/book-meeting-room/internal/domain"
)

type FindUserRepo struct {
	Database DatabaseOperator
}

func NewFindUserRepo(db DatabaseOperator) *FindUserRepo {
	return &FindUserRepo{Database: db}
}

func (f FindUserRepo) FindUser(ctx context.Context, username string) (domain.User, error) {
	var dbUser []User
	query := `SELECT id, username, password FROM users WHERE username = ?`
	err := f.Database.Read(ctx, &dbUser, query, username)
	if err != nil {
		return domain.User{}, err
	}

	if len(dbUser) < 1 {
		return domain.User{}, sql.ErrNoRows
	}

	foundUser := dbUser[0]
	user := domain.User{
		ID:       foundUser.ID,
		Username: foundUser.Username,
		Password: foundUser.Password,
	}
	return user, nil
}
