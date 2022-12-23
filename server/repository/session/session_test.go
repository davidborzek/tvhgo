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

// Test data models
var testUser = &core.User{
	Username:    "testuser",
	Email:       "testuser@example.com",
	DisplayName: "Test user",
}

func initTestUser(db *sql.DB) error {
	return user.New(db).
		Create(noCtx, testUser)
}

func TestMain(m *testing.M) {
	db, err := testdb.Setup()
	if err != nil {
		panic(err)
	}
	defer testdb.Close(db)

	if err := initTestUser(db); err != nil {
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

func TestCreate(t *testing.T) {
	session := &core.Session{
		UserId:      testUser.ID,
		HashedToken: "someHashedToken",
		ClientIP:    "127.0.0.1",
		UserAgent:   "someUserAgent",
		CreatedAt:   time.Now().Unix(),
		LastUsedAt:  time.Now().Unix(),
	}

	err := repository.Create(noCtx, session)
	assert.Nil(t, err)
	assert.NotEmpty(t, session.ID)

	t.Run("Find", testFind(session))
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
		err := repository.Delete(noCtx, created.ID)

		assert.Nil(t, err)

		session, err := repository.Find(noCtx, created.HashedToken)

		assert.Nil(t, err)
		assert.Nil(t, session)
	}
}
