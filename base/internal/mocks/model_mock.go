// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coderbiq/pointsgo/base/internal/model (interfaces: AccountRepository,AccountLogStorer,Infra,AppServices,RegisterService,DepositService,ConsumeService,AccountFinder)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	devent "github.com/coderbiq/dgo/base/devent"
	vo "github.com/coderbiq/dgo/base/vo"
	model "github.com/coderbiq/pointsgo/base/internal/model"
	common "github.com/coderbiq/pointsgo/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountRepository is a mock of AccountRepository interface
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// FindByOwner mocks base method
func (m *MockAccountRepository) FindByOwner(arg0 vo.LongID) ([]model.Account, error) {
	ret := m.ctrl.Call(m, "FindByOwner", arg0)
	ret0, _ := ret[0].([]model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOwner indicates an expected call of FindByOwner
func (mr *MockAccountRepositoryMockRecorder) FindByOwner(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOwner", reflect.TypeOf((*MockAccountRepository)(nil).FindByOwner), arg0)
}

// Get mocks base method
func (m *MockAccountRepository) Get(arg0 vo.LongID) (model.Account, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAccountRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountRepository)(nil).Get), arg0)
}

// Save mocks base method
func (m *MockAccountRepository) Save(arg0 model.Account) error {
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockAccountRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccountRepository)(nil).Save), arg0)
}

// MockAccountLogStorer is a mock of AccountLogStorer interface
type MockAccountLogStorer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountLogStorerMockRecorder
}

// MockAccountLogStorerMockRecorder is the mock recorder for MockAccountLogStorer
type MockAccountLogStorerMockRecorder struct {
	mock *MockAccountLogStorer
}

// NewMockAccountLogStorer creates a new mock instance
func NewMockAccountLogStorer(ctrl *gomock.Controller) *MockAccountLogStorer {
	mock := &MockAccountLogStorer{ctrl: ctrl}
	mock.recorder = &MockAccountLogStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountLogStorer) EXPECT() *MockAccountLogStorerMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockAccountLogStorer) Append(arg0 common.AccountLog) {
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append
func (mr *MockAccountLogStorerMockRecorder) Append(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockAccountLogStorer)(nil).Append), arg0)
}

// Get mocks base method
func (m *MockAccountLogStorer) Get(arg0 vo.Identity) []common.AccountLog {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]common.AccountLog)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockAccountLogStorerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountLogStorer)(nil).Get), arg0)
}

// MockInfra is a mock of Infra interface
type MockInfra struct {
	ctrl     *gomock.Controller
	recorder *MockInfraMockRecorder
}

// MockInfraMockRecorder is the mock recorder for MockInfra
type MockInfraMockRecorder struct {
	mock *MockInfra
}

// NewMockInfra creates a new mock instance
func NewMockInfra(ctrl *gomock.Controller) *MockInfra {
	mock := &MockInfra{ctrl: ctrl}
	mock.recorder = &MockInfraMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfra) EXPECT() *MockInfraMockRecorder {
	return m.recorder
}

// AccountRepo mocks base method
func (m *MockInfra) AccountRepo() model.AccountRepository {
	ret := m.ctrl.Call(m, "AccountRepo")
	ret0, _ := ret[0].(model.AccountRepository)
	return ret0
}

// AccountRepo indicates an expected call of AccountRepo
func (mr *MockInfraMockRecorder) AccountRepo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountRepo", reflect.TypeOf((*MockInfra)(nil).AccountRepo))
}

// EventBus mocks base method
func (m *MockInfra) EventBus() devent.Bus {
	ret := m.ctrl.Call(m, "EventBus")
	ret0, _ := ret[0].(devent.Bus)
	return ret0
}

// EventBus indicates an expected call of EventBus
func (mr *MockInfraMockRecorder) EventBus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventBus", reflect.TypeOf((*MockInfra)(nil).EventBus))
}

// LogStorer mocks base method
func (m *MockInfra) LogStorer() model.AccountLogStorer {
	ret := m.ctrl.Call(m, "LogStorer")
	ret0, _ := ret[0].(model.AccountLogStorer)
	return ret0
}

// LogStorer indicates an expected call of LogStorer
func (mr *MockInfraMockRecorder) LogStorer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogStorer", reflect.TypeOf((*MockInfra)(nil).LogStorer))
}

// MockAppServices is a mock of AppServices interface
type MockAppServices struct {
	ctrl     *gomock.Controller
	recorder *MockAppServicesMockRecorder
}

// MockAppServicesMockRecorder is the mock recorder for MockAppServices
type MockAppServicesMockRecorder struct {
	mock *MockAppServices
}

// NewMockAppServices creates a new mock instance
func NewMockAppServices(ctrl *gomock.Controller) *MockAppServices {
	mock := &MockAppServices{ctrl: ctrl}
	mock.recorder = &MockAppServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppServices) EXPECT() *MockAppServicesMockRecorder {
	return m.recorder
}

// ConsumeApp mocks base method
func (m *MockAppServices) ConsumeApp() model.ConsumeService {
	ret := m.ctrl.Call(m, "ConsumeApp")
	ret0, _ := ret[0].(model.ConsumeService)
	return ret0
}

// ConsumeApp indicates an expected call of ConsumeApp
func (mr *MockAppServicesMockRecorder) ConsumeApp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeApp", reflect.TypeOf((*MockAppServices)(nil).ConsumeApp))
}

// DepositApp mocks base method
func (m *MockAppServices) DepositApp() model.DepositService {
	ret := m.ctrl.Call(m, "DepositApp")
	ret0, _ := ret[0].(model.DepositService)
	return ret0
}

// DepositApp indicates an expected call of DepositApp
func (mr *MockAppServicesMockRecorder) DepositApp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositApp", reflect.TypeOf((*MockAppServices)(nil).DepositApp))
}

// Finder mocks base method
func (m *MockAppServices) Finder() model.AccountFinder {
	ret := m.ctrl.Call(m, "Finder")
	ret0, _ := ret[0].(model.AccountFinder)
	return ret0
}

// Finder indicates an expected call of Finder
func (mr *MockAppServicesMockRecorder) Finder() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finder", reflect.TypeOf((*MockAppServices)(nil).Finder))
}

// RegisterApp mocks base method
func (m *MockAppServices) RegisterApp() model.RegisterService {
	ret := m.ctrl.Call(m, "RegisterApp")
	ret0, _ := ret[0].(model.RegisterService)
	return ret0
}

// RegisterApp indicates an expected call of RegisterApp
func (mr *MockAppServicesMockRecorder) RegisterApp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterApp", reflect.TypeOf((*MockAppServices)(nil).RegisterApp))
}

// RunTasks mocks base method
func (m *MockAppServices) RunTasks(arg0 context.Context) {
	m.ctrl.Call(m, "RunTasks", arg0)
}

// RunTasks indicates an expected call of RunTasks
func (mr *MockAppServicesMockRecorder) RunTasks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTasks", reflect.TypeOf((*MockAppServices)(nil).RunTasks), arg0)
}

// MockRegisterService is a mock of RegisterService interface
type MockRegisterService struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterServiceMockRecorder
}

// MockRegisterServiceMockRecorder is the mock recorder for MockRegisterService
type MockRegisterServiceMockRecorder struct {
	mock *MockRegisterService
}

// NewMockRegisterService creates a new mock instance
func NewMockRegisterService(ctrl *gomock.Controller) *MockRegisterService {
	mock := &MockRegisterService{ctrl: ctrl}
	mock.recorder = &MockRegisterServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterService) EXPECT() *MockRegisterServiceMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockRegisterService) Register(arg0 string) (int64, error) {
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockRegisterServiceMockRecorder) Register(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegisterService)(nil).Register), arg0)
}

// MockDepositService is a mock of DepositService interface
type MockDepositService struct {
	ctrl     *gomock.Controller
	recorder *MockDepositServiceMockRecorder
}

// MockDepositServiceMockRecorder is the mock recorder for MockDepositService
type MockDepositServiceMockRecorder struct {
	mock *MockDepositService
}

// NewMockDepositService creates a new mock instance
func NewMockDepositService(ctrl *gomock.Controller) *MockDepositService {
	mock := &MockDepositService{ctrl: ctrl}
	mock.recorder = &MockDepositServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDepositService) EXPECT() *MockDepositServiceMockRecorder {
	return m.recorder
}

// Deposit mocks base method
func (m *MockDepositService) Deposit(arg0 int64, arg1 uint) (uint, uint, error) {
	ret := m.ctrl.Call(m, "Deposit", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Deposit indicates an expected call of Deposit
func (mr *MockDepositServiceMockRecorder) Deposit(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockDepositService)(nil).Deposit), arg0, arg1)
}

// MockConsumeService is a mock of ConsumeService interface
type MockConsumeService struct {
	ctrl     *gomock.Controller
	recorder *MockConsumeServiceMockRecorder
}

// MockConsumeServiceMockRecorder is the mock recorder for MockConsumeService
type MockConsumeServiceMockRecorder struct {
	mock *MockConsumeService
}

// NewMockConsumeService creates a new mock instance
func NewMockConsumeService(ctrl *gomock.Controller) *MockConsumeService {
	mock := &MockConsumeService{ctrl: ctrl}
	mock.recorder = &MockConsumeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumeService) EXPECT() *MockConsumeServiceMockRecorder {
	return m.recorder
}

// Consume mocks base method
func (m *MockConsumeService) Consume(arg0 int64, arg1 uint) (uint, uint, error) {
	ret := m.ctrl.Call(m, "Consume", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Consume indicates an expected call of Consume
func (mr *MockConsumeServiceMockRecorder) Consume(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockConsumeService)(nil).Consume), arg0, arg1)
}

// MockAccountFinder is a mock of AccountFinder interface
type MockAccountFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAccountFinderMockRecorder
}

// MockAccountFinderMockRecorder is the mock recorder for MockAccountFinder
type MockAccountFinderMockRecorder struct {
	mock *MockAccountFinder
}

// NewMockAccountFinder creates a new mock instance
func NewMockAccountFinder(ctrl *gomock.Controller) *MockAccountFinder {
	mock := &MockAccountFinder{ctrl: ctrl}
	mock.recorder = &MockAccountFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountFinder) EXPECT() *MockAccountFinderMockRecorder {
	return m.recorder
}

// Detail mocks base method
func (m *MockAccountFinder) Detail(arg0 int64) (common.AccountReader, error) {
	ret := m.ctrl.Call(m, "Detail", arg0)
	ret0, _ := ret[0].(common.AccountReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail
func (mr *MockAccountFinderMockRecorder) Detail(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockAccountFinder)(nil).Detail), arg0)
}
