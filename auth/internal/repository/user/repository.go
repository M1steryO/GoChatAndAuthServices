package user

import (
	"auth/internal/repository"
	modelRepo "auth/internal/repository/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user exists")
)

type repo struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) repository.UserRepository {
	return &repo{
		pool: pool,
	}
}

func (s *repo) Get(ctx context.Context, id int64) (*modelRepo.User, error) {
	user := modelRepo.User{}
	var role string
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, name, created_at, updated_at, role
			 FROM "user"
			 WHERE id=$1`,
		id).Scan(&user.Id, &user.Info.Email,
		&user.Info.Name, &user.CreatedAt,
		&user.UpdatedAt, &role)
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
	err := s.pool.QueryRow(ctx,
		`INSERT INTO "user" (email, name, password, role) 
			 VALUES ($1, $2, $3, $4) 
			 RETURNING id;`,
		user.Info.Email, user.Info.Name, user.Password, user.Info.Role).Scan(&lastInsertId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return 0, ErrUserExists
			}
		}
		return 0, err
	}
	return lastInsertId, nil

}
