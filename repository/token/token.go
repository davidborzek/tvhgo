package token

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidborzek/tvhgo/core"
)

type sqlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) core.TokenRepository {
	return &sqlRepository{db: db}
}

func (s *sqlRepository) FindByToken(ctx context.Context, hashedToken string) (*core.Token, error) {
	row := s.db.QueryRowContext(ctx, queryByToken, sql.Named("hashed_token", hashedToken))

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
	rows, err := s.db.QueryContext(ctx, queryByUser, sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

func (s *sqlRepository) Create(ctx context.Context, token *core.Token) error {
	now := time.Now().Unix()

	res, err := s.db.ExecContext(ctx, stmtInsert,
		sql.Named("user_id", token.UserID),
		sql.Named("name", token.Name),
		sql.Named("hashed_token", token.HashedToken),
		sql.Named("created_at", now),
		sql.Named("updated_at", now),
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

func (s *sqlRepository) Delete(ctx context.Context, token *core.Token) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, sql.Named("id", token.ID))
	return err
}
