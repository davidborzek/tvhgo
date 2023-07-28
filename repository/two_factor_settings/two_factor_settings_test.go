package twofactorsettings_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/db/testdb"
	twofactorsettings "github.com/davidborzek/tvhgo/repository/two_factor_settings"
	"github.com/davidborzek/tvhgo/repository/user"
	"github.com/davidborzek/tvhgo/services/clock"
	"github.com/stretchr/testify/assert"
)

var (
	noCtx      = context.TODO()
	repository core.TwoFactorSettingsRepository

	testUser = &core.User{
		Username:    "testuser",
		Email:       "testuser@example.com",
		DisplayName: "Test user",
	}
)

func initTestUser(db *sql.DB) error {
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

	repository = twofactorsettings.New(db)
	code := m.Run()

	err = testdb.TruncateTables(db, "two_factor_settings", "user")
	if err != nil {
		panic(err)
	}

	testdb.Close(db)

	os.Exit(code)
}

func TestFindReturnsNil(t *testing.T) {
	settings, err := repository.Find(noCtx, 0)

	assert.Nil(t, settings)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	settings := &core.TwoFactorSettings{
		UserID:  testUser.ID,
		Secret:  "someSecret",
		Enabled: true,
	}
	err := repository.Create(noCtx, settings)

	assert.Nil(t, err)

	t.Run("Find", testFind(settings))
	t.Run("Delete", testDelete(settings))
}

func testFind(created *core.TwoFactorSettings) func(t *testing.T) {
	return func(t *testing.T) {
		settings, err := repository.Find(noCtx, created.UserID)

		assert.Nil(t, err)
		assert.Equal(t, created, settings)
	}
}

func testDelete(created *core.TwoFactorSettings) func(t *testing.T) {
	return func(t *testing.T) {
		err := repository.Delete(noCtx, created)

		assert.Nil(t, err)

		settings, err := repository.Find(noCtx, created.UserID)

		assert.Nil(t, err)
		assert.Nil(t, settings)
	}
}
