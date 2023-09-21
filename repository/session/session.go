package session

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/davidborzek/tvhgo/core"
)

type sqlRepository struct {
	db    *sql.DB
	clock core.Clock
}

func New(db *sql.DB, clock core.Clock) core.SessionRepository {
	return &sqlRepository{
		db:    db,
		clock: clock,
	}
}

func (s *sqlRepository) Find(ctx context.Context, hashedToken string) (*core.Session, error) {
	row := s.db.QueryRowContext(ctx, queryByToken, hashedToken)
	session := new(core.Session)
	if err := scanRow(row, session); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to query session row: %w", err)
	}

	return session, nil
}

func (s *sqlRepository) FindByUser(ctx context.Context, userID int64) ([]*core.Session, error) {
	rows, err := s.db.QueryContext(ctx, queryByUserID, userID)
	if err != nil {
		return nil, fmt.Errorf("session FindByUser query failed: %w", err)
	}

	return scanRows(rows)
}

func (s *sqlRepository) Create(ctx context.Context, session *core.Session) error {
	now := s.clock.Now().Unix()

	res, err := s.db.ExecContext(ctx, stmtInsert,
		session.UserId,
		session.HashedToken,
		session.ClientIP,
		session.UserAgent,
		now,
		now,
		now,
	)

	if err != nil {
		return fmt.Errorf("failed to exec session insert: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed retrieve last insert id for session: %w", err)
	}

	session.ID = id
	session.CreatedAt = now
	session.LastUsedAt = now
	session.RotatedAt = now
	return nil
}

func (s *sqlRepository) Update(ctx context.Context, session *core.Session) error {
	_, err := s.db.ExecContext(ctx, stmtUpdate,
		session.HashedToken,
		session.ClientIP,
		session.UserAgent,
		session.LastUsedAt,
		session.RotatedAt,
		session.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to exec session update: %w", err)
	}

	return nil
}

func (s *sqlRepository) Delete(ctx context.Context, sessionID int64, userID int64) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, sessionID, userID)

	if err != nil {
		return fmt.Errorf("failed to exec session delete: %w", err)
	}

	return nil
}

func (s *sqlRepository) DeleteExpired(
	ctx context.Context,
	expirationDate int64,
	inactiveExpirationDate int64,
) (int64, error) {
	res, err := s.db.ExecContext(ctx, stmtDeleteExpired, expirationDate, inactiveExpirationDate)
	if err != nil {
		return 0, fmt.Errorf("failed to exec session delete expired: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed get rows affected count for session delete expired: %w", err)
	}

	return rows, nil
}
