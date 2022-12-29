package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidborzek/tvhgo/core"
)

type sqlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) core.UserRepository {
	return &sqlRepository{
		db: db,
	}
}

func (s *sqlRepository) FindById(ctx context.Context, id int64) (*core.User, error) {
	row := s.db.QueryRowContext(ctx, queryById, id)
	user := new(core.User)
	if err := scanRow(row, user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return user, nil
}

func (s *sqlRepository) FindByUsername(ctx context.Context, username string) (*core.User, error) {
	row := s.db.QueryRowContext(ctx, queryByUsername, username)
	user := new(core.User)
	if err := scanRow(row, user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (s *sqlRepository) Find(ctx context.Context, params core.UserQueryParams) ([]*core.User, error) {
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

	return scanRows(rows)
}

func (s *sqlRepository) Create(ctx context.Context, user *core.User) error {
	createdAt := time.Now().Unix()

	res, err := s.db.ExecContext(ctx, stmtInsert,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.DisplayName,
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

func (s *sqlRepository) Update(ctx context.Context, user *core.User) error {
	updatedAt := time.Now().Unix()

	_, err := s.db.ExecContext(ctx, stmtUpdate,
		user.Username,
		user.PasswordHash,
		user.Email,
		user.DisplayName,
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
