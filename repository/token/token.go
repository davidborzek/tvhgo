package token

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db"
)

type sqlRepository struct {
	db *db.DB
}

func New(db *db.DB) core.TokenRepository {
	return &sqlRepository{db: db}
}

func (s *sqlRepository) FindByToken(ctx context.Context, hashedToken string) (*core.Token, error) {
	row := s.db.QueryRowContext(ctx, queryByToken, hashedToken)

	token := new(core.Token)
	if err := scanRow(row, token); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return token, nil
}

func (s *sqlRepository) FindByUser(ctx context.Context, userID int64) ([]*core.Token, error) {
	rows, err := s.db.QueryContext(ctx, queryByUser, userID)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

func (s *sqlRepository) Create(ctx context.Context, token *core.Token) error {
	if s.db.Type == config.DatabaseTypePostgres {
		return s.createPostgres(ctx, token)
	}

	return s.create(ctx, token)
}

func (s *sqlRepository) create(ctx context.Context, token *core.Token) error {
	now := time.Now().Unix()

	res, err := s.db.ExecContext(ctx, stmtInsert,
		token.UserID,
		token.HashedToken,
		token.Name,
		now,
		now,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	token.ID = id
	token.UpdatedAt = now
	token.CreatedAt = now
	return nil
}

func (s *sqlRepository) createPostgres(ctx context.Context, token *core.Token) error {
	now := time.Now().Unix()

	err := s.db.QueryRowContext(ctx, stmtInsertPostgres,
		token.UserID,
		token.HashedToken,
		token.Name,
		now,
		now,
	).Scan(&token.ID)

	if err != nil {
		return err
	}

	token.UpdatedAt = now
	token.CreatedAt = now
	return nil
}

func (s *sqlRepository) Delete(ctx context.Context, token *core.Token) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, token.ID)
	return err
}
