package token_test

import (
	"context"
	"os"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	database "github.com/davidborzek/tvhgo/db"
	"github.com/davidborzek/tvhgo/db/testdb"
	"github.com/davidborzek/tvhgo/repository/token"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/stretchr/testify/assert"
)

var (
	noCtx      = context.TODO()
	repository core.TokenRepository

	testUser = &core.User{
		ID:          1234,
		Username:    "testuser",
		Email:       "testuser@example.com",
		DisplayName: "Test user",
	}
)

func initTestUser(db *database.DB) error {
	return user.New(db, clock.NewClock()).
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

	repository = token.New(db)
	code := m.Run()

	err = testdb.TruncateTables(db, "token", "user")
	if err != nil {
		panic(err)
	}

	testdb.Close(db)

	os.Exit(code)
}

func TestFindReturnsNil(t *testing.T) {
	token, err := repository.FindByToken(noCtx, "unknownToken")

	assert.Nil(t, token)
	assert.Nil(t, err)
}

func TestFindByUserReturnsEmptyArray(t *testing.T) {
	tokens, err := repository.FindByUser(noCtx, testUser.ID)

	assert.Empty(t, tokens)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	token := &core.Token{
		UserID:      testUser.ID,
		HashedToken: "someToken",
		Name:        "someName",
	}
	err := repository.Create(noCtx, token)

	assert.Nil(t, err)

	t.Run("Find", testFind(token))
	t.Run("FindByUser", testFindByUser(token))
	t.Run("Delete", testDelete(token))
}

func testFind(created *core.Token) func(t *testing.T) {
	return func(t *testing.T) {
		token, err := repository.FindByToken(noCtx, created.HashedToken)

		assert.Nil(t, err)
		assert.Equal(t, created, token)
	}
}

func testFindByUser(created *core.Token) func(t *testing.T) {
	return func(t *testing.T) {
		tokens, err := repository.FindByUser(noCtx, testUser.ID)

		assert.Nil(t, err)
		assert.Len(t, tokens, 1)
		assert.Equal(t, created, tokens[0])
	}
}

func testDelete(created *core.Token) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Delete(noCtx, created)

		assert.Nil(t, err)

		settings, err := repository.FindByToken(noCtx, created.HashedToken)

		assert.Nil(t, err)
		assert.Nil(t, settings)
	}
}
