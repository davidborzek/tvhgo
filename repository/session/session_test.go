package session_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db/testdb"
	"github.com/davidborzek/tvhgo/repository/session"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/stretchr/testify/assert"
)

var noCtx = context.TODO()

var repository core.SessionRepository

var db *sql.DB

// Test data models
var testUser = &core.User{
	Username:    "testuser",
	Email:       "testuser@example.com",
	DisplayName: "Test user",
}

func initTestUser() error {
	return user.New(db).
		Create(noCtx, testUser)
}

func TestMain(m *testing.M) {
	_db, err := testdb.Setup()
	if err != nil {
		panic(err)
	}
	db = _db
	defer testdb.Close(db)

	if err := initTestUser(); err != nil {
		panic(err)
	}

	repository = session.New(db)
	code := m.Run()

	err = testdb.TruncateTables(db, "session", "user")
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestFindReturnsNil(t *testing.T) {
	session, err := repository.Find(noCtx, "unknown")

	assert.Nil(t, session)
	assert.Nil(t, err)
}

func TestFindByUserReturnsEmptyArray(t *testing.T) {
	sessions, err := repository.FindByUser(noCtx, 0)

	assert.Empty(t, sessions)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	session := &core.Session{
		UserId:      testUser.ID,
		HashedToken: "someHashedToken",
		ClientIP:    "127.0.0.1",
		UserAgent:   "someUserAgent",
	}

	err := repository.Create(noCtx, session)
	assert.Nil(t, err)
	assert.NotEmpty(t, session.ID)
	assert.NotEmpty(t, session.CreatedAt)
	assert.NotEmpty(t, session.LastUsedAt)
	assert.NotEmpty(t, session.RotatedAt)

	t.Run("Find", testFind(session))
	t.Run("FindByUser", testFindByUser(session))
	t.Run("Update", testUpdate(session))
	t.Run("Delete", testDelete(session))
}

func testFind(created *core.Session) func(t *testing.T) {
	return func(t *testing.T) {
		session, err := repository.Find(noCtx, created.HashedToken)

		assert.Nil(t, err)
		assert.Equal(t, created, session)
	}
}

func testFindByUser(created *core.Session) func(t *testing.T) {
	return func(t *testing.T) {
		sessions, err := repository.FindByUser(noCtx, created.UserId)

		assert.Nil(t, err)
		assert.Len(t, sessions, 1)
		assert.Equal(t, created, sessions[0])
	}
}

func testUpdate(created *core.Session) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Update(noCtx, &core.Session{
			ID:          created.ID,
			HashedToken: "updatedHashedToken",
			ClientIP:    "10.0.0.1",
			UserAgent:   "updatedUserAgent",
			LastUsedAt:  int64(9999),
		})

		assert.Nil(t, err)

		session, err := repository.Find(noCtx, "updatedHashedToken")
		assert.Nil(t, err)

		assert.Equal(t, "10.0.0.1", session.ClientIP)
		assert.Equal(t, "updatedUserAgent", session.UserAgent)
		assert.Equal(t, int64(9999), session.LastUsedAt)
	}
}

func testDelete(created *core.Session) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Delete(noCtx, created.ID, created.UserId)

		assert.Nil(t, err)

		session, err := repository.Find(noCtx, created.HashedToken)

		assert.Nil(t, err)
		assert.Nil(t, session)
	}
}

func TestDeleteExpired(t *testing.T) {
	q := `INSERT INTO session
	(user_id, hashed_token, client_ip, user_agent, created_at, last_used_at, rotated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now().Unix()

	_, err := db.Exec(q, testUser.ID, "token1", "", "", 111, now, 0)
	assert.Nil(t, err)

	_, err = db.Exec(q, testUser.ID, "token2", "", "", now, 111, 0)
	assert.Nil(t, err)

	_, err = db.Exec(q, testUser.ID, "token3", "", "", now, now, 0)
	assert.Nil(t, err)

	rows, err := repository.DeleteExpired(noCtx, 222, 222)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), rows)
}
