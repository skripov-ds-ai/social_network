package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/skripov-ds-ai/social_network/internal/models"
)

type Repository interface {
	Get(ctx context.Context, userID string) (models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Update(ctx context.Context, user models.UpdateUserParams) error
	Delete(ctx context.Context, userID string) error
}

type storage struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &storage{
		db: db,
	}
}

func (s *storage) Get(ctx context.Context, userID string) (models.User, error) {
	user := models.User{}
	args := map[string]any{
		"id": userID,
	}
	q := `
select
	*
from public.users
where id = :id`
	res, err := s.db.NamedQuery(q, args)
	if err != nil {
		return models.User{}, err
	}
	ok := res.Next()
	if !ok {
		return models.User{}, models.ErrNotFound
	}
	err = res.StructScan(&user)
	if err != nil {
		return models.User{}, err
	}
	err = res.Err()
	return user, err
}

func (s *storage) Create(ctx context.Context, user models.User) (string, error) {
	args := map[string]any{
		"first_name":    user.FirstName,
		"second_name":   user.SecondName,
		"birthdate":     user.Birthdate,
		"biography":     user.Biography,
		"city":          user.City,
		"password_hash": user.PasswordHash,
	}
	q := `
insert into public.users
    (first_name, second_name, birthdate, biography, city, password_hash)
values
	(:first_name, :second_name, :birthdate, :biography, :city, :password_hash)
returning id`
	res, err := s.db.NamedQuery(q, args)
	if err != nil {
		return "", err
	}
	var userID string
	for res.Next() {
		err = res.Scan(&userID)
		if err != nil {
			return "", err
		}
	}
	err = res.Err()
	return userID, err
}

func (s *storage) Update(ctx context.Context, user models.UpdateUserParams) error {
	return nil
}

func (s *storage) Delete(ctx context.Context, userID string) error {
	return nil
}
