package user

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	modelRepo "auth/internal/repository/user/model"
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user exists")
)

const constraintErrorCode = "23505"

type repo struct {
	db db.Client
}

func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (s *repo) Get(ctx context.Context, id int64) (*modelRepo.User, error) {
	user := modelRepo.User{}
	q := db.Query{
		Title: "user_repository.Get",
		Query: `SELECT id, username, name, created_at, updated_at, role
			 FROM "user"
			 WHERE id=$1`,
	}
	err := s.db.DB().ScanOneContext(ctx, &user, q, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *repo) Create(ctx context.Context, user *modelRepo.User) (int64, error) {
	var lastInsertId int64
	q := db.Query{
		Title: "user_repository.Create",
		Query: `INSERT INTO "user" (username, name, password, role) 
				VALUES ($1, $2, $3, $4) 
			 	RETURNING id;`,
	}
	err := s.db.DB().QueryRowContext(ctx, q,
		user.Info.Username, user.Info.Name, user.Password, user.Info.Role).Scan(&lastInsertId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == constraintErrorCode {
				return 0, ErrUserExists
			}
		}
		return 0, err
	}
	return lastInsertId, nil

}

func (s *repo) CreateLog(ctx context.Context, userId int64, action string) error {
	q := db.Query{
		Title: "user_repository.CreateLog",
		Query: `INSERT INTO "user_log" (user_id, action)
				VALUES ($1, $2)`,
	}

	_, err := s.db.DB().ExecContext(ctx, q, userId, action)
	if err != nil {
		return err
	}
	return nil
}
