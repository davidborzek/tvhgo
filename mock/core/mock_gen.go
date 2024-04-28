// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/davidborzek/tvhgo/core (interfaces: UserRepository,SessionRepository,Clock,TwoFactorAuthService,TwoFactorSettingsRepository,TokenRepository,TokenService,SessionManager,ChannelService)
//
// Generated by this command:
//
//	mockgen -destination=mock_gen.go github.com/davidborzek/tvhgo/core UserRepository,SessionRepository,Clock,TwoFactorAuthService,TwoFactorSettingsRepository,TokenRepository,TokenService,SessionManager,ChannelService
//

// Package mock_core is a generated GoMock package.
package mock_core

import (
	context "context"
	reflect "reflect"
	time "time"

	core "github.com/davidborzek/tvhgo/core"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockUserRepository) Delete(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepository)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockUserRepository) Find(arg0 context.Context, arg1 core.UserQueryParams) (*core.UserListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.UserListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserRepositoryMockRecorder) Find(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepository)(nil).Find), arg0, arg1)
}

// FindById mocks base method.
func (m *MockUserRepository) FindById(arg0 context.Context, arg1 int64) (*core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserRepositoryMockRecorder) FindById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserRepository)(nil).FindById), arg0, arg1)
}

// FindByUsername mocks base method.
func (m *MockUserRepository) FindByUsername(arg0 context.Context, arg1 string) (*core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", arg0, arg1)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockUserRepositoryMockRecorder) FindByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindByUsername), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserRepository) Update(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0, arg1)
}

// MockSessionRepository is a mock of SessionRepository interface.
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository.
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance.
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionRepository) Create(arg0 context.Context, arg1 *core.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSessionRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSessionRepository) Delete(arg0 context.Context, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSessionRepositoryMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionRepository)(nil).Delete), arg0, arg1, arg2)
}

// DeleteExpired mocks base method.
func (m *MockSessionRepository) DeleteExpired(arg0 context.Context, arg1, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpired", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpired indicates an expected call of DeleteExpired.
func (mr *MockSessionRepositoryMockRecorder) DeleteExpired(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockSessionRepository)(nil).DeleteExpired), arg0, arg1, arg2)
}

// Find mocks base method.
func (m *MockSessionRepository) Find(arg0 context.Context, arg1 string) (*core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockSessionRepositoryMockRecorder) Find(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSessionRepository)(nil).Find), arg0, arg1)
}

// FindByUser mocks base method.
func (m *MockSessionRepository) FindByUser(arg0 context.Context, arg1 int64) ([]*core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0, arg1)
	ret0, _ := ret[0].([]*core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser.
func (mr *MockSessionRepositoryMockRecorder) FindByUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockSessionRepository)(nil).FindByUser), arg0, arg1)
}

// Update mocks base method.
func (m *MockSessionRepository) Update(arg0 context.Context, arg1 *core.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSessionRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSessionRepository)(nil).Update), arg0, arg1)
}

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// Now mocks base method.
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// MockTwoFactorAuthService is a mock of TwoFactorAuthService interface.
type MockTwoFactorAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockTwoFactorAuthServiceMockRecorder
}

// MockTwoFactorAuthServiceMockRecorder is the mock recorder for MockTwoFactorAuthService.
type MockTwoFactorAuthServiceMockRecorder struct {
	mock *MockTwoFactorAuthService
}

// NewMockTwoFactorAuthService creates a new mock instance.
func NewMockTwoFactorAuthService(ctrl *gomock.Controller) *MockTwoFactorAuthService {
	mock := &MockTwoFactorAuthService{ctrl: ctrl}
	mock.recorder = &MockTwoFactorAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwoFactorAuthService) EXPECT() *MockTwoFactorAuthServiceMockRecorder {
	return m.recorder
}

// Activate mocks base method.
func (m *MockTwoFactorAuthService) Activate(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Activate indicates an expected call of Activate.
func (mr *MockTwoFactorAuthServiceMockRecorder) Activate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockTwoFactorAuthService)(nil).Activate), arg0, arg1, arg2)
}

// Deactivate mocks base method.
func (m *MockTwoFactorAuthService) Deactivate(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockTwoFactorAuthServiceMockRecorder) Deactivate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockTwoFactorAuthService)(nil).Deactivate), arg0, arg1, arg2)
}

// GetSettings mocks base method.
func (m *MockTwoFactorAuthService) GetSettings(arg0 context.Context, arg1 int64) (*core.TwoFactorSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", arg0, arg1)
	ret0, _ := ret[0].(*core.TwoFactorSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockTwoFactorAuthServiceMockRecorder) GetSettings(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockTwoFactorAuthService)(nil).GetSettings), arg0, arg1)
}

// Setup mocks base method.
func (m *MockTwoFactorAuthService) Setup(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Setup indicates an expected call of Setup.
func (mr *MockTwoFactorAuthServiceMockRecorder) Setup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockTwoFactorAuthService)(nil).Setup), arg0, arg1)
}

// Verify mocks base method.
func (m *MockTwoFactorAuthService) Verify(arg0 context.Context, arg1 int64, arg2 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockTwoFactorAuthServiceMockRecorder) Verify(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTwoFactorAuthService)(nil).Verify), arg0, arg1, arg2)
}

// MockTwoFactorSettingsRepository is a mock of TwoFactorSettingsRepository interface.
type MockTwoFactorSettingsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTwoFactorSettingsRepositoryMockRecorder
}

// MockTwoFactorSettingsRepositoryMockRecorder is the mock recorder for MockTwoFactorSettingsRepository.
type MockTwoFactorSettingsRepositoryMockRecorder struct {
	mock *MockTwoFactorSettingsRepository
}

// NewMockTwoFactorSettingsRepository creates a new mock instance.
func NewMockTwoFactorSettingsRepository(ctrl *gomock.Controller) *MockTwoFactorSettingsRepository {
	mock := &MockTwoFactorSettingsRepository{ctrl: ctrl}
	mock.recorder = &MockTwoFactorSettingsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwoFactorSettingsRepository) EXPECT() *MockTwoFactorSettingsRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTwoFactorSettingsRepository) Create(arg0 context.Context, arg1 *core.TwoFactorSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTwoFactorSettingsRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTwoFactorSettingsRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTwoFactorSettingsRepository) Delete(arg0 context.Context, arg1 *core.TwoFactorSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTwoFactorSettingsRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTwoFactorSettingsRepository)(nil).Delete), arg0, arg1)
}

// Find mocks base method.
func (m *MockTwoFactorSettingsRepository) Find(arg0 context.Context, arg1 int64) (*core.TwoFactorSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.TwoFactorSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockTwoFactorSettingsRepositoryMockRecorder) Find(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockTwoFactorSettingsRepository)(nil).Find), arg0, arg1)
}

// Save mocks base method.
func (m *MockTwoFactorSettingsRepository) Save(arg0 context.Context, arg1 *core.TwoFactorSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTwoFactorSettingsRepositoryMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTwoFactorSettingsRepository)(nil).Save), arg0, arg1)
}

// Update mocks base method.
func (m *MockTwoFactorSettingsRepository) Update(arg0 context.Context, arg1 *core.TwoFactorSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTwoFactorSettingsRepositoryMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTwoFactorSettingsRepository)(nil).Update), arg0, arg1)
}

// MockTokenRepository is a mock of TokenRepository interface.
type MockTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRepositoryMockRecorder
}

// MockTokenRepositoryMockRecorder is the mock recorder for MockTokenRepository.
type MockTokenRepositoryMockRecorder struct {
	mock *MockTokenRepository
}

// NewMockTokenRepository creates a new mock instance.
func NewMockTokenRepository(ctrl *gomock.Controller) *MockTokenRepository {
	mock := &MockTokenRepository{ctrl: ctrl}
	mock.recorder = &MockTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRepository) EXPECT() *MockTokenRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTokenRepository) Create(arg0 context.Context, arg1 *core.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTokenRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTokenRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTokenRepository) Delete(arg0 context.Context, arg1 *core.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTokenRepositoryMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTokenRepository)(nil).Delete), arg0, arg1)
}

// FindByToken mocks base method.
func (m *MockTokenRepository) FindByToken(arg0 context.Context, arg1 string) (*core.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByToken", arg0, arg1)
	ret0, _ := ret[0].(*core.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByToken indicates an expected call of FindByToken.
func (mr *MockTokenRepositoryMockRecorder) FindByToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByToken", reflect.TypeOf((*MockTokenRepository)(nil).FindByToken), arg0, arg1)
}

// FindByUser mocks base method.
func (m *MockTokenRepository) FindByUser(arg0 context.Context, arg1 int64) ([]*core.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0, arg1)
	ret0, _ := ret[0].([]*core.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser.
func (mr *MockTokenRepositoryMockRecorder) FindByUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockTokenRepository)(nil).FindByUser), arg0, arg1)
}

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTokenService) Create(arg0 context.Context, arg1 int64, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTokenServiceMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTokenService)(nil).Create), arg0, arg1, arg2)
}

// Revoke mocks base method.
func (m *MockTokenService) Revoke(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke.
func (mr *MockTokenServiceMockRecorder) Revoke(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockTokenService)(nil).Revoke), arg0, arg1)
}

// Validate mocks base method.
func (m *MockTokenService) Validate(arg0 context.Context, arg1 string) (*core.AuthContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(*core.AuthContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenServiceMockRecorder) Validate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenService)(nil).Validate), arg0, arg1)
}

// MockSessionManager is a mock of SessionManager interface.
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager.
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance.
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionManager) Create(arg0 context.Context, arg1 int64, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSessionManagerMockRecorder) Create(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionManager)(nil).Create), arg0, arg1, arg2, arg3)
}

// Revoke mocks base method.
func (m *MockSessionManager) Revoke(arg0 context.Context, arg1, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke.
func (mr *MockSessionManagerMockRecorder) Revoke(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockSessionManager)(nil).Revoke), arg0, arg1, arg2)
}

// Validate mocks base method.
func (m *MockSessionManager) Validate(arg0 context.Context, arg1 string) (*core.AuthContext, *string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(*core.AuthContext)
	ret1, _ := ret[1].(*string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Validate indicates an expected call of Validate.
func (mr *MockSessionManagerMockRecorder) Validate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockSessionManager)(nil).Validate), arg0, arg1)
}

// MockChannelService is a mock of ChannelService interface.
type MockChannelService struct {
	ctrl     *gomock.Controller
	recorder *MockChannelServiceMockRecorder
}

// MockChannelServiceMockRecorder is the mock recorder for MockChannelService.
type MockChannelServiceMockRecorder struct {
	mock *MockChannelService
}

// NewMockChannelService creates a new mock instance.
func NewMockChannelService(ctrl *gomock.Controller) *MockChannelService {
	mock := &MockChannelService{ctrl: ctrl}
	mock.recorder = &MockChannelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelService) EXPECT() *MockChannelServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockChannelService) Get(arg0 context.Context, arg1 string) (*core.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*core.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockChannelServiceMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChannelService)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockChannelService) GetAll(arg0 context.Context, arg1 core.PaginationSortQueryParams) ([]*core.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]*core.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockChannelServiceMockRecorder) GetAll(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockChannelService)(nil).GetAll), arg0, arg1)
}
