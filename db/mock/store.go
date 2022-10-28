// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/peacewalker122/project/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/peacewalker122/project/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccounts mocks base method.
func (m *MockStore) CreateAccounts(arg0 context.Context, arg1 db.CreateAccountsParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccounts", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccounts indicates an expected call of CreateAccounts.
func (mr *MockStoreMockRecorder) CreateAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccounts", reflect.TypeOf((*MockStore)(nil).CreateAccounts), arg0, arg1)
}

// CreateComment_feature mocks base method.
func (m *MockStore) CreateComment_feature(arg0 context.Context, arg1 db.CreateComment_featureParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment_feature", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment_feature indicates an expected call of CreateComment_feature.
func (mr *MockStoreMockRecorder) CreateComment_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment_feature", reflect.TypeOf((*MockStore)(nil).CreateComment_feature), arg0, arg1)
}

// CreateEntries mocks base method.
func (m *MockStore) CreateEntries(arg0 context.Context, arg1 db.CreateEntriesParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntries indicates an expected call of CreateEntries.
func (mr *MockStoreMockRecorder) CreateEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntries", reflect.TypeOf((*MockStore)(nil).CreateEntries), arg0, arg1)
}

// CreateLike_feature mocks base method.
func (m *MockStore) CreateLike_feature(arg0 context.Context, arg1 db.CreateLike_featureParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike_feature", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLike_feature indicates an expected call of CreateLike_feature.
func (mr *MockStoreMockRecorder) CreateLike_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike_feature", reflect.TypeOf((*MockStore)(nil).CreateLike_feature), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreatePost_feature mocks base method.
func (m *MockStore) CreatePost_feature(arg0 context.Context, arg1 int64) (db.PostFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost_feature", arg0, arg1)
	ret0, _ := ret[0].(db.PostFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost_feature indicates an expected call of CreatePost_feature.
func (mr *MockStoreMockRecorder) CreatePost_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost_feature", reflect.TypeOf((*MockStore)(nil).CreatePost_feature), arg0, arg1)
}

// CreateQouteRetweet_feature mocks base method.
func (m *MockStore) CreateQouteRetweet_feature(arg0 context.Context, arg1 db.CreateQouteRetweet_featureParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQouteRetweet_feature", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQouteRetweet_feature indicates an expected call of CreateQouteRetweet_feature.
func (mr *MockStoreMockRecorder) CreateQouteRetweet_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQouteRetweet_feature", reflect.TypeOf((*MockStore)(nil).CreateQouteRetweet_feature), arg0, arg1)
}

// CreateRetweet_feature mocks base method.
func (m *MockStore) CreateRetweet_feature(arg0 context.Context, arg1 db.CreateRetweet_featureParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRetweet_feature", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRetweet_feature indicates an expected call of CreateRetweet_feature.
func (mr *MockStoreMockRecorder) CreateRetweet_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRetweet_feature", reflect.TypeOf((*MockStore)(nil).CreateRetweet_feature), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockStore) DeletePost(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockStoreMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockStore)(nil).DeletePost), arg0, arg1)
}

// DeletePostFeature mocks base method.
func (m *MockStore) DeletePostFeature(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostFeature", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostFeature indicates an expected call of DeletePostFeature.
func (mr *MockStoreMockRecorder) DeletePostFeature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostFeature", reflect.TypeOf((*MockStore)(nil).DeletePostFeature), arg0, arg1)
}

// DeleteQouteRetweet mocks base method.
func (m *MockStore) DeleteQouteRetweet(arg0 context.Context, arg1 db.DeleteQouteRetweetParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQouteRetweet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQouteRetweet indicates an expected call of DeleteQouteRetweet.
func (mr *MockStoreMockRecorder) DeleteQouteRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQouteRetweet", reflect.TypeOf((*MockStore)(nil).DeleteQouteRetweet), arg0, arg1)
}

// GetAccounts mocks base method.
func (m *MockStore) GetAccounts(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockStoreMockRecorder) GetAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockStore)(nil).GetAccounts), arg0, arg1)
}

// GetAccountsOwner mocks base method.
func (m *MockStore) GetAccountsOwner(arg0 context.Context, arg1 string) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsOwner", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsOwner indicates an expected call of GetAccountsOwner.
func (mr *MockStoreMockRecorder) GetAccountsOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsOwner", reflect.TypeOf((*MockStore)(nil).GetAccountsOwner), arg0, arg1)
}

// GetEntries mocks base method.
func (m *MockStore) GetEntries(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockStoreMockRecorder) GetEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockStore)(nil).GetEntries), arg0, arg1)
}

// GetEntriesFull mocks base method.
func (m *MockStore) GetEntriesFull(arg0 context.Context, arg1 db.GetEntriesFullParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntriesFull", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetEntriesFull indicates an expected call of GetEntriesFull.
func (mr *MockStoreMockRecorder) GetEntriesFull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntriesFull", reflect.TypeOf((*MockStore)(nil).GetEntriesFull), arg0, arg1)
}

// GetLikeInfo mocks base method.
func (m *MockStore) GetLikeInfo(arg0 context.Context, arg1 db.GetLikeInfoParams) (db.LikeFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeInfo", arg0, arg1)
	ret0, _ := ret[0].(db.LikeFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeInfo indicates an expected call of GetLikeInfo.
func (mr *MockStoreMockRecorder) GetLikeInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeInfo", reflect.TypeOf((*MockStore)(nil).GetLikeInfo), arg0, arg1)
}

// GetLikejoin mocks base method.
func (m *MockStore) GetLikejoin(arg0 context.Context, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikejoin", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikejoin indicates an expected call of GetLikejoin.
func (mr *MockStoreMockRecorder) GetLikejoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikejoin", reflect.TypeOf((*MockStore)(nil).GetLikejoin), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockStore) GetPost(arg0 context.Context, arg1 int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockStoreMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStore)(nil).GetPost), arg0, arg1)
}

// GetPostInfoJoin mocks base method.
func (m *MockStore) GetPostInfoJoin(arg0 context.Context, arg1 db.GetPostInfoJoinParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostInfoJoin", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostInfoJoin indicates an expected call of GetPostInfoJoin.
func (mr *MockStoreMockRecorder) GetPostInfoJoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostInfoJoin", reflect.TypeOf((*MockStore)(nil).GetPostInfoJoin), arg0, arg1)
}

// GetPostJoin mocks base method.
func (m *MockStore) GetPostJoin(arg0 context.Context, arg1 int64) (db.GetPostJoinRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostJoin", arg0, arg1)
	ret0, _ := ret[0].(db.GetPostJoinRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostJoin indicates an expected call of GetPostJoin.
func (mr *MockStoreMockRecorder) GetPostJoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostJoin", reflect.TypeOf((*MockStore)(nil).GetPostJoin), arg0, arg1)
}

// GetPostJoin_QouteRetweet mocks base method.
func (m *MockStore) GetPostJoin_QouteRetweet(arg0 context.Context, arg1 db.GetPostJoin_QouteRetweetParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostJoin_QouteRetweet", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostJoin_QouteRetweet indicates an expected call of GetPostJoin_QouteRetweet.
func (mr *MockStoreMockRecorder) GetPostJoin_QouteRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostJoin_QouteRetweet", reflect.TypeOf((*MockStore)(nil).GetPostJoin_QouteRetweet), arg0, arg1)
}

// GetPost_feature mocks base method.
func (m *MockStore) GetPost_feature(arg0 context.Context, arg1 int64) (db.PostFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost_feature", arg0, arg1)
	ret0, _ := ret[0].(db.PostFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost_feature indicates an expected call of GetPost_feature.
func (mr *MockStoreMockRecorder) GetPost_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost_feature", reflect.TypeOf((*MockStore)(nil).GetPost_feature), arg0, arg1)
}

// GetPost_feature_Update mocks base method.
func (m *MockStore) GetPost_feature_Update(arg0 context.Context, arg1 int64) (db.PostFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost_feature_Update", arg0, arg1)
	ret0, _ := ret[0].(db.PostFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost_feature_Update indicates an expected call of GetPost_feature_Update.
func (mr *MockStoreMockRecorder) GetPost_feature_Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost_feature_Update", reflect.TypeOf((*MockStore)(nil).GetPost_feature_Update), arg0, arg1)
}

// GetQouteRetweet mocks base method.
func (m *MockStore) GetQouteRetweet(arg0 context.Context, arg1 db.GetQouteRetweetParams) (db.QouteRetweetFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQouteRetweet", arg0, arg1)
	ret0, _ := ret[0].(db.QouteRetweetFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQouteRetweet indicates an expected call of GetQouteRetweet.
func (mr *MockStoreMockRecorder) GetQouteRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQouteRetweet", reflect.TypeOf((*MockStore)(nil).GetQouteRetweet), arg0, arg1)
}

// GetQouteRetweetJoin mocks base method.
func (m *MockStore) GetQouteRetweetJoin(arg0 context.Context, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQouteRetweetJoin", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQouteRetweetJoin indicates an expected call of GetQouteRetweetJoin.
func (mr *MockStoreMockRecorder) GetQouteRetweetJoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQouteRetweetJoin", reflect.TypeOf((*MockStore)(nil).GetQouteRetweetJoin), arg0, arg1)
}

// GetQouteRetweetRows mocks base method.
func (m *MockStore) GetQouteRetweetRows(arg0 context.Context, arg1 db.GetQouteRetweetRowsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQouteRetweetRows", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQouteRetweetRows indicates an expected call of GetQouteRetweetRows.
func (mr *MockStoreMockRecorder) GetQouteRetweetRows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQouteRetweetRows", reflect.TypeOf((*MockStore)(nil).GetQouteRetweetRows), arg0, arg1)
}

// GetRetweet mocks base method.
func (m *MockStore) GetRetweet(arg0 context.Context, arg1 db.GetRetweetParams) (db.RetweetFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetweet", arg0, arg1)
	ret0, _ := ret[0].(db.RetweetFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRetweet indicates an expected call of GetRetweet.
func (mr *MockStoreMockRecorder) GetRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetweet", reflect.TypeOf((*MockStore)(nil).GetRetweet), arg0, arg1)
}

// GetRetweetJoin mocks base method.
func (m *MockStore) GetRetweetJoin(arg0 context.Context, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetweetJoin", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRetweetJoin indicates an expected call of GetRetweetJoin.
func (mr *MockStoreMockRecorder) GetRetweetJoin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetweetJoin", reflect.TypeOf((*MockStore)(nil).GetRetweetJoin), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSessionuser mocks base method.
func (m *MockStore) GetSessionuser(arg0 context.Context, arg1 string) (db.GetSessionuserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionuser", arg0, arg1)
	ret0, _ := ret[0].(db.GetSessionuserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionuser indicates an expected call of GetSessionuser.
func (mr *MockStoreMockRecorder) GetSessionuser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionuser", reflect.TypeOf((*MockStore)(nil).GetSessionuser), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListComment mocks base method.
func (m *MockStore) ListComment(arg0 context.Context, arg1 db.ListCommentParams) ([]db.ListCommentRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComment", arg0, arg1)
	ret0, _ := ret[0].([]db.ListCommentRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComment indicates an expected call of ListComment.
func (mr *MockStoreMockRecorder) ListComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComment", reflect.TypeOf((*MockStore)(nil).ListComment), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListPost mocks base method.
func (m *MockStore) ListPost(arg0 context.Context, arg1 db.ListPostParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPost", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPost indicates an expected call of ListPost.
func (mr *MockStoreMockRecorder) ListPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPost", reflect.TypeOf((*MockStore)(nil).ListPost), arg0, arg1)
}

// ListUser mocks base method.
func (m *MockStore) ListUser(arg0 context.Context, arg1 db.ListUserParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockStoreMockRecorder) ListUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockStore)(nil).ListUser), arg0, arg1)
}

// UpdateLike mocks base method.
func (m *MockStore) UpdateLike(arg0 context.Context, arg1 db.UpdateLikeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLike indicates an expected call of UpdateLike.
func (mr *MockStoreMockRecorder) UpdateLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLike", reflect.TypeOf((*MockStore)(nil).UpdateLike), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), arg0, arg1)
}

// UpdatePost_feature mocks base method.
func (m *MockStore) UpdatePost_feature(arg0 context.Context, arg1 db.UpdatePost_featureParams) (db.PostFeature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost_feature", arg0, arg1)
	ret0, _ := ret[0].(db.PostFeature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost_feature indicates an expected call of UpdatePost_feature.
func (mr *MockStoreMockRecorder) UpdatePost_feature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost_feature", reflect.TypeOf((*MockStore)(nil).UpdatePost_feature), arg0, arg1)
}

// UpdateQouteRetweet mocks base method.
func (m *MockStore) UpdateQouteRetweet(arg0 context.Context, arg1 db.UpdateQouteRetweetParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQouteRetweet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQouteRetweet indicates an expected call of UpdateQouteRetweet.
func (mr *MockStoreMockRecorder) UpdateQouteRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQouteRetweet", reflect.TypeOf((*MockStore)(nil).UpdateQouteRetweet), arg0, arg1)
}

// UpdateRetweet mocks base method.
func (m *MockStore) UpdateRetweet(arg0 context.Context, arg1 db.UpdateRetweetParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRetweet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRetweet indicates an expected call of UpdateRetweet.
func (mr *MockStoreMockRecorder) UpdateRetweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRetweet", reflect.TypeOf((*MockStore)(nil).UpdateRetweet), arg0, arg1)
}
