package session

import (
	"context"
	"database/sql"
	"time"

	"github.com/davidborzek/tvhgo/core"
)

type sqlRepository struct {
	db *sql.DB
}

func New(db *sql.DB) core.SessionRepository {
	return &sqlRepository{
		db: db,
	}
}

func (s *sqlRepository) Find(ctx context.Context, hashedToken string) (*core.Session, error) {
	row := s.db.QueryRowContext(ctx, queryByToken, hashedToken)
	session := new(core.Session)
	if err := scanRow(row, session); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return session, nil
}

func (s *sqlRepository) FindByUser(ctx context.Context, userID int64) ([]*core.Session, error) {
	rows, err := s.db.QueryContext(ctx, queryByUserID, userID)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

func (s *sqlRepository) Create(ctx context.Context, session *core.Session) error {
	now := time.Now().Unix()

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
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
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

	return err
}

func (s *sqlRepository) Delete(ctx context.Context, sessionID int64, userID int64) error {
	_, err := s.db.ExecContext(ctx, stmtDelete, sessionID, userID)
	return err
}

func (s *sqlRepository) DeleteExpired(ctx context.Context, expirationDate int64, inactiveExpirationDate int64) (int64, error) {
	res, err := s.db.ExecContext(ctx, stmtDeleteExpired, expirationDate, inactiveExpirationDate)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
