package repository

import (
	utils "auth/internal/utils/storage"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user exists")
)

type User struct {
	Id        int64
	Email     string
	Name      string
	Role      int
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
type Storage struct {
	Pool *pgxpool.Pool
}

func (s *Storage) GetUser(ctx context.Context, id int64) (*User, error) {
	user := User{}
	var role string
	err := s.Pool.QueryRow(ctx,
		`SELECT id, email, name, created_at, updated_at, role
			 FROM "user"
			 WHERE id=$1`,
		id).Scan(&user.Id, &user.Email,
		&user.Name, &user.CreatedAt,
		&user.UpdatedAt, &role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	user.Role = utils.GetRoleIdByStr(role)
	return &user, nil
}

func (s *Storage) CreateUser(ctx context.Context, user User) (int64, error) {
	var lastInsertId int64
	err := s.Pool.QueryRow(ctx,
		`INSERT INTO "user" (email, name, password, role) 
			 VALUES ($1, $2, $3, $4) 
			 RETURNING id;`,
		user.Email, user.Name, user.Password, user.Role).Scan(&lastInsertId)
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
