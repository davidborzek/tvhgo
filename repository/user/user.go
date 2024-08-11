package user

import (
	"context"
	"database/sql"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db"
)

type sqlRepository struct {
	db    *db.DB
	clock core.Clock
}

func New(db *db.DB, clock core.Clock) core.UserRepository {
	return &sqlRepository{
		db:    db,
		clock: clock,
	}
}

func (s *sqlRepository) FindById(ctx context.Context, id int64) (*core.User, error) {
	return s.findBy(ctx, queryById, id)
}

func (s *sqlRepository) FindByUsername(ctx context.Context, username string) (*core.User, error) {
	return s.findBy(ctx, queryByUsername, username)

}

func (s *sqlRepository) FindByEmail(ctx context.Context, email string) (*core.User, error) {
	return s.findBy(ctx, queryByEmail, email)
}

func (s *sqlRepository) Find(
	ctx context.Context,
	params core.UserQueryParams,
) (*core.UserListResult, error) {
	var count int64
	err := s.db.QueryRowContext(ctx, queryCount).Scan(&count)
	if err != nil {
		return nil, err
	}

	args := []interface{}{}
	query := queryBase
	if params.Limit > 0 {
		query += ` LIMIT ?`
		args = append(args, params.Limit)
	}

	if params.Limit > 0 && params.Offset > 0 {
		query += ` OFFSET ?`
		args = append(args, params.Offset)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	users, err := scanRows(rows)
	if err != nil {
		return nil, err
	}

	return &core.UserListResult{
		Entries: users,
		Total:   count,
		Offset:  params.Offset,
	}, nil
}

func (s *sqlRepository) Create(ctx context.Context, user *core.User) error {
	if err := s.validateUserUnique(ctx, user); err != nil {
		return err
	}

	if s.db.Type == config.DatabaseTypePostgres {
		return s.createPostgres(ctx, user)
	}

	return s.create(ctx, user)
}

func (s *sqlRepository) create(ctx context.Context, user *core.User) error {
	createdAt := s.clock.Now().Unix()

	res, err := s.db.ExecContext(ctx, stmtInsert,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.DisplayName,
		user.IsAdmin,
		createdAt,
		createdAt,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	user.UpdatedAt = createdAt
	user.CreatedAt = createdAt

	return nil
}

func (s *sqlRepository) createPostgres(ctx context.Context, user *core.User) error {
	createdAt := s.clock.Now().Unix()

	err := s.db.QueryRowContext(ctx, stmtInsertPostgres,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.DisplayName,
		user.IsAdmin,
		createdAt,
		createdAt,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	user.UpdatedAt = createdAt
	user.CreatedAt = createdAt

	return nil
}

func (s *sqlRepository) Update(ctx context.Context, user *core.User) error {
	if err := s.validateUserUnique(ctx, user); err != nil {
		return err
	}

	updatedAt := s.clock.Now().Unix()

	_, err := s.db.ExecContext(ctx, stmtUpdate,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.DisplayName,
		user.IsAdmin,
		updatedAt,
		user.ID,
	)

	if err == nil {
		user.UpdatedAt = updatedAt
	}

	return err
}

func (s *sqlRepository) Delete(ctx context.Context, user *core.User) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, user.ID)
	return err
}

func (s *sqlRepository) findBy(
	ctx context.Context,
	query string,
	args ...interface{},
) (*core.User, error) {
	row := s.db.QueryRowContext(ctx, query, args...)
	user := new(core.User)
	if err := scanRow(row, user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return user, nil
}

func (s *sqlRepository) validateUserUnique(ctx context.Context, user *core.User) error {
	if err := s.validateUsernameUnique(ctx, user); err != nil {
		return err
	}

	if err := s.validateEmailUnique(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *sqlRepository) validateUsernameUnique(ctx context.Context, user *core.User) error {
	maybeUser, err := s.FindByUsername(ctx, user.Username)
	if err != nil {
		return err
	}

	if maybeUser == nil {
		return nil
	}

	if maybeUser.ID == user.ID {
		return nil
	}

	return core.ErrUsernameAlreadyExists
}

func (s *sqlRepository) validateEmailUnique(ctx context.Context, user *core.User) error {
	maybeUser, err := s.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if maybeUser == nil {
		return nil
	}

	if maybeUser.ID == user.ID {
		return nil
	}

	return core.ErrEmailAlreadyExists
}
