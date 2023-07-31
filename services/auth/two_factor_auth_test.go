package auth_test

import (
	"errors"
	"testing"
	"time"

	"github.com/davidborzek/tvhgo/config"
	"github.com/davidborzek/tvhgo/core"
	mock_core "github.com/davidborzek/tvhgo/mock/core"
	"github.com/davidborzek/tvhgo/services/auth"
	"github.com/golang/mock/gomock"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

const (
	totpSecret = "VR52N6Y5EK7DAEBUHAPMJ3KSEDWYMCHC"
)

var cfg = &config.TOTPConfig{
	Issuer: "tvhgo",
}

func TestTwoFactorServiceVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	code, err := totp.GenerateCode(totpSecret, time.Now())
	if err != nil {
		assert.Nil(t, err)
		return
	}

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(&core.TwoFactorSettings{
			Secret: totpSecret,
		}, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err = twoFactorService.Verify(ctx, userID, &code)
	assert.Nil(t, err)
}

func TestTwoFactorServiceVerifyWhenTwoFactorIsNotEnabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(nil, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err := twoFactorService.Verify(ctx, userID, nil)
	assert.Nil(t, err)
}

func TestTwoFactorServiceVerifyReturnsErrTwoFactorRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)
	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(&core.TwoFactorSettings{
			Enabled: true,
		}, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err := twoFactorService.Verify(ctx, userID, nil)
	assert.ErrorIs(t, err, core.ErrTwoFactorRequired)
}

func TestTwoFactorServiceVerifyReturnsErrTwoFactorCodeInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(&core.TwoFactorSettings{
			Enabled: true,
			Secret:  totpSecret,
		}, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	code := "invalid"

	err := twoFactorService.Verify(ctx, userID, &code)
	assert.ErrorIs(t, err, core.ErrTwoFactorCodeInvalid)
}

func TestGetSettingsReturnsTwoFactorSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	expectedSettings := &core.TwoFactorSettings{
		Enabled: true,
		Secret:  totpSecret,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(expectedSettings, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	settings, err := twoFactorService.GetSettings(ctx, userID)
	assert.Equal(t, expectedSettings, settings)
	assert.Nil(t, err)
}

func TestGetSettingsReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(nil, errors.New("some error")).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	settings, err := twoFactorService.GetSettings(ctx, userID)
	assert.Nil(t, settings)
	assert.NotNil(t, err)
}

func TestGetSettingsReturnsDefaultWhenNoSettingsArePresent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(nil, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	settings, err := twoFactorService.GetSettings(ctx, userID)
	assert.False(t, settings.Enabled)
	assert.Nil(t, err)
}

func TestSetupGeneratesAndPersistTOTPSecret(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	mockUserRepository.EXPECT().
		FindById(ctx, userID).
		Return(&core.User{
			ID:       userID,
			Username: username,
		}, nil).
		Times(1)

	mockTwoFactorSettingsRepository.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	url, err := twoFactorService.Setup(ctx, userID)
	assert.Contains(t, url, "otpauth://")
	assert.Contains(t, url, username)
	assert.Nil(t, err)
}

func TestDeactivate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	settings := &core.TwoFactorSettings{
		UserID: userID,
		Secret: totpSecret,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(settings, nil).
		Times(1)

	mockTwoFactorSettingsRepository.EXPECT().
		Delete(ctx, settings).
		Return(nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	code, err := totp.GenerateCode(totpSecret, time.Now())
	if err != nil {
		assert.Nil(t, err)
		return
	}

	err = twoFactorService.Deactivate(ctx, userID, code)
	assert.Nil(t, err)
}

func TestDeactivateReturnsErrTwoFactorCodeInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	settings := &core.TwoFactorSettings{
		UserID: userID,
		Secret: totpSecret,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(settings, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err := twoFactorService.Deactivate(ctx, userID, "invalid")
	assert.ErrorIs(t, err, core.ErrTwoFactorCodeInvalid)
}

func TestActivate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	settings := &core.TwoFactorSettings{
		UserID: userID,
		Secret: totpSecret,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(settings, nil).
		Times(1)

	mockTwoFactorSettingsRepository.EXPECT().
		Update(ctx, settings).
		Return(nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	code, err := totp.GenerateCode(totpSecret, time.Now())
	if err != nil {
		assert.Nil(t, err)
		return
	}

	err = twoFactorService.Activate(ctx, userID, code)
	assert.Nil(t, err)
}

func TestActivateReturnsErrTwoFactorCodeInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	settings := &core.TwoFactorSettings{
		UserID: userID,
		Secret: totpSecret,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(settings, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err := twoFactorService.Activate(ctx, userID, "invalid")
	assert.ErrorIs(t, err, core.ErrTwoFactorCodeInvalid)
}

func TestActivateReturnsErrTwoFactorAuthAlreadyEnabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_core.NewMockUserRepository(ctrl)
	mockTwoFactorSettingsRepository := mock_core.NewMockTwoFactorSettingsRepository(ctrl)

	settings := &core.TwoFactorSettings{
		UserID:  userID,
		Secret:  totpSecret,
		Enabled: true,
	}

	mockTwoFactorSettingsRepository.EXPECT().
		Find(ctx, userID).
		Return(settings, nil).
		Times(1)

	twoFactorService := auth.NewTwoFactorAuthService(
		mockTwoFactorSettingsRepository,
		mockUserRepository,
		cfg,
	)

	err := twoFactorService.Activate(ctx, userID, "ignored")
	assert.ErrorIs(t, err, core.ErrTwoFactorAuthAlreadyEnabled)
}
