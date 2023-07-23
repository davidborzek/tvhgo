package user_test

import (
	"context"
	"os"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db/testdb"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/stretchr/testify/assert"
)

var noCtx = context.TODO()

var repository core.UserRepository

func TestMain(m *testing.M) {
	db, err := testdb.Setup()
	if err != nil {
		panic(err)
	}

	repository = user.New(db)
	code := m.Run()

	err = testdb.TruncateTables(db, "user")
	if err != nil {
		panic(err)
	}

	testdb.Close(db)

	os.Exit(code)
}

func TestFindByIdReturnsNil(t *testing.T) {
	user, err := repository.FindById(noCtx, 1)

	assert.Nil(t, user)
	assert.Nil(t, err)
}

func TestFindByUsernameReturnsNil(t *testing.T) {
	user, err := repository.FindByUsername(noCtx, "unknown-username")

	assert.Nil(t, user)
	assert.Nil(t, err)
}

func TestFindReturnsEmptyArray(t *testing.T) {
	users, err := repository.Find(noCtx, core.UserQueryParams{})

	assert.Nil(t, err)
	assert.Empty(t, users)
}

func TestCreate(t *testing.T) {
	user := &core.User{
		Username:     "test_user",
		PasswordHash: "somePasswordHash",
		Email:        "test@example.com",
		DisplayName:  "Test User",
	}
	err := repository.Create(noCtx, user)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.ID)

	t.Run("FindById", testFindById(user))
	t.Run("FindByUsername", testFindByUsername(user))
	t.Run("Find", testFind(user))
	t.Run("FindOffset", testFindOffset(user))
	t.Run("CreateWithExistingUsername", testCreateWithExistingUsername(user))
	t.Run("CreateWithExistingMail", testCreateWithExistingEmail(user))
	t.Run("Update", testUpdate(user))
	t.Run("Delete", testDelete(user))
}

func testFindById(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := repository.FindById(noCtx, created.ID)

		assert.Nil(t, err)
		assert.Equal(t, created, user)
	}
}

func testFindByUsername(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := repository.FindByUsername(noCtx, created.Username)

		assert.Nil(t, err)
		assert.Equal(t, created, user)
	}
}

func testFind(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		users, err := repository.Find(noCtx, core.UserQueryParams{
			Limit: 1,
		})

		assert.Nil(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, created, users[0])
	}
}

func testFindOffset(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		users, err := repository.Find(noCtx, core.UserQueryParams{
			Limit:  1,
			Offset: 1,
		})

		assert.Nil(t, err)
		assert.Empty(t, users)
	}
}

func testCreateWithExistingUsername(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Create(noCtx, &core.User{
			Username: created.Username,
		})
		assert.Equal(t, core.ErrUsernameAlreadyExists, err)
	}
}

func testCreateWithExistingEmail(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Create(noCtx, &core.User{
			Username: "newUser",
			Email:    created.Email,
		})
		assert.Equal(t, core.ErrEmailAlreadyExists, err)
	}
}

func testUpdate(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Update(noCtx,
			&core.User{
				ID:           created.ID,
				Username:     "updated_user",
				PasswordHash: "updatedPasswordHash",
				DisplayName:  "Updated User",
				Email:        "updated@example.com",
			},
		)

		assert.Nil(t, err)

		user, err := repository.FindById(noCtx, created.ID)
		assert.Nil(t, err)

		assert.Equal(t, "updated_user", user.Username)
		assert.Equal(t, "updatedPasswordHash", user.PasswordHash)
		assert.Equal(t, "Updated User", user.DisplayName)
		assert.Equal(t, "updated@example.com", user.Email)

		assert.NotEqual(t, 0, user.UpdatedAt)
	}
}

func testDelete(created *core.User) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Delete(noCtx, &core.User{ID: created.ID})
		assert.Nil(t, err)

		user, err := repository.FindById(noCtx, created.ID)
		assert.Nil(t, user)
		assert.Nil(t, err)
	}
}
