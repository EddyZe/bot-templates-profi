package userrepo

import (
	"bot-templates-profi/internal/domain/entity"
	"bot-templates-profi/internal/storage/postgres"
	"context"
	"errors"
	"log"
)

const (
	UserIsExist = "user is exist"
)

type UserRepo interface {
	Save(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context) []entity.User
}

type UserRepository struct {
	*postgres.Postgres
}

func New(psql *postgres.Postgres) *UserRepository {
	return &UserRepository{
		psql,
	}
}

func (r *UserRepository) Save(ctx context.Context, usr *entity.User) error {
	query, args, err := r.BindNamed(
		`insert into usr (telegram_id, username) values (:telegram_id, :username)`,
		usr,
	)
	if err != nil {
		return err
	}

	if err := r.QueryRowxContext(
		ctx,
		query,
		args...).Err(); err != nil {
		return errors.New(UserIsExist)
	}

	return nil
}

func (r UserRepository) FindAll(ctx context.Context) []entity.User {
	res := make([]entity.User, 0)

	if err := r.SelectContext(
		ctx,
		&res,
		`select * from usr`,
	); err != nil {
		log.Println(err)
	}

	return res
}
