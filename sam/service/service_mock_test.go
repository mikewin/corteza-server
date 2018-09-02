// Code generated by MockGen. DO NOT EDIT.
// Source: sam/service/service_test.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	repository "github.com/crusttech/crust/sam/repository"
	types "github.com/crusttech/crust/sam/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// BeginWith mocks base method
func (m *MockRepository) BeginWith(ctx context.Context, callback repository.BeginCallback) error {
	ret := m.ctrl.Call(m, "BeginWith", ctx, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginWith indicates an expected call of BeginWith
func (mr *MockRepositoryMockRecorder) BeginWith(ctx, callback interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginWith", reflect.TypeOf((*MockRepository)(nil).BeginWith), ctx, callback)
}

// Begin mocks base method
func (m *MockRepository) Begin() error {
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin
func (mr *MockRepositoryMockRecorder) Begin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockRepository)(nil).Begin))
}

// Rollback mocks base method
func (m *MockRepository) Rollback() error {
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockRepositoryMockRecorder) Rollback() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockRepository)(nil).Rollback))
}

// Commit mocks base method
func (m *MockRepository) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockRepositoryMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockRepository)(nil).Commit))
}

// WithCtx mocks base method
func (m *MockRepository) WithCtx(ctx context.Context) repository.Interfaces {
	ret := m.ctrl.Call(m, "WithCtx", ctx)
	ret0, _ := ret[0].(repository.Interfaces)
	return ret0
}

// WithCtx indicates an expected call of WithCtx
func (mr *MockRepositoryMockRecorder) WithCtx(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCtx", reflect.TypeOf((*MockRepository)(nil).WithCtx), ctx)
}

// FindAttachmentByID mocks base method
func (m *MockRepository) FindAttachmentByID(id uint64) (*types.Attachment, error) {
	ret := m.ctrl.Call(m, "FindAttachmentByID", id)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAttachmentByID indicates an expected call of FindAttachmentByID
func (mr *MockRepositoryMockRecorder) FindAttachmentByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAttachmentByID", reflect.TypeOf((*MockRepository)(nil).FindAttachmentByID), id)
}

// FindAttachmentByRange mocks base method
func (m *MockRepository) FindAttachmentByRange(channelID, fromAttachmentID, toAttachmentID uint64) ([]*types.Attachment, error) {
	ret := m.ctrl.Call(m, "FindAttachmentByRange", channelID, fromAttachmentID, toAttachmentID)
	ret0, _ := ret[0].([]*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAttachmentByRange indicates an expected call of FindAttachmentByRange
func (mr *MockRepositoryMockRecorder) FindAttachmentByRange(channelID, fromAttachmentID, toAttachmentID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAttachmentByRange", reflect.TypeOf((*MockRepository)(nil).FindAttachmentByRange), channelID, fromAttachmentID, toAttachmentID)
}

// CreateAttachment mocks base method
func (m *MockRepository) CreateAttachment(mod *types.Attachment) (*types.Attachment, error) {
	ret := m.ctrl.Call(m, "CreateAttachment", mod)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttachment indicates an expected call of CreateAttachment
func (mr *MockRepositoryMockRecorder) CreateAttachment(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttachment", reflect.TypeOf((*MockRepository)(nil).CreateAttachment), mod)
}

// UpdateAttachment mocks base method
func (m *MockRepository) UpdateAttachment(mod *types.Attachment) (*types.Attachment, error) {
	ret := m.ctrl.Call(m, "UpdateAttachment", mod)
	ret0, _ := ret[0].(*types.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttachment indicates an expected call of UpdateAttachment
func (mr *MockRepositoryMockRecorder) UpdateAttachment(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttachment", reflect.TypeOf((*MockRepository)(nil).UpdateAttachment), mod)
}

// DeleteAttachmentByID mocks base method
func (m *MockRepository) DeleteAttachmentByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteAttachmentByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAttachmentByID indicates an expected call of DeleteAttachmentByID
func (mr *MockRepositoryMockRecorder) DeleteAttachmentByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttachmentByID", reflect.TypeOf((*MockRepository)(nil).DeleteAttachmentByID), id)
}

// FindChannelByID mocks base method
func (m *MockRepository) FindChannelByID(id uint64) (*types.Channel, error) {
	ret := m.ctrl.Call(m, "FindChannelByID", id)
	ret0, _ := ret[0].(*types.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChannelByID indicates an expected call of FindChannelByID
func (mr *MockRepositoryMockRecorder) FindChannelByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChannelByID", reflect.TypeOf((*MockRepository)(nil).FindChannelByID), id)
}

// FindDirectChannelByUserID mocks base method
func (m *MockRepository) FindDirectChannelByUserID(fromUserID, toUserID uint64) (*types.Channel, error) {
	ret := m.ctrl.Call(m, "FindDirectChannelByUserID", fromUserID, toUserID)
	ret0, _ := ret[0].(*types.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDirectChannelByUserID indicates an expected call of FindDirectChannelByUserID
func (mr *MockRepositoryMockRecorder) FindDirectChannelByUserID(fromUserID, toUserID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDirectChannelByUserID", reflect.TypeOf((*MockRepository)(nil).FindDirectChannelByUserID), fromUserID, toUserID)
}

// FindChannels mocks base method
func (m *MockRepository) FindChannels(filter *types.ChannelFilter) ([]*types.Channel, error) {
	ret := m.ctrl.Call(m, "FindChannels", filter)
	ret0, _ := ret[0].([]*types.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChannels indicates an expected call of FindChannels
func (mr *MockRepositoryMockRecorder) FindChannels(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChannels", reflect.TypeOf((*MockRepository)(nil).FindChannels), filter)
}

// CreateChannel mocks base method
func (m *MockRepository) CreateChannel(mod *types.Channel) (*types.Channel, error) {
	ret := m.ctrl.Call(m, "CreateChannel", mod)
	ret0, _ := ret[0].(*types.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChannel indicates an expected call of CreateChannel
func (mr *MockRepositoryMockRecorder) CreateChannel(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannel", reflect.TypeOf((*MockRepository)(nil).CreateChannel), mod)
}

// UpdateChannel mocks base method
func (m *MockRepository) UpdateChannel(mod *types.Channel) (*types.Channel, error) {
	ret := m.ctrl.Call(m, "UpdateChannel", mod)
	ret0, _ := ret[0].(*types.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChannel indicates an expected call of UpdateChannel
func (mr *MockRepositoryMockRecorder) UpdateChannel(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChannel", reflect.TypeOf((*MockRepository)(nil).UpdateChannel), mod)
}

// FindChannelsMembershipsByMemberId mocks base method
func (m *MockRepository) FindChannelsMembershipsByMemberId(memberId uint64) ([]*types.ChannelMember, error) {
	ret := m.ctrl.Call(m, "FindChannelsMembershipsByMemberId", memberId)
	ret0, _ := ret[0].([]*types.ChannelMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindChannelsMembershipsByMemberId indicates an expected call of FindChannelsMembershipsByMemberId
func (mr *MockRepositoryMockRecorder) FindChannelsMembershipsByMemberId(memberId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindChannelsMembershipsByMemberId", reflect.TypeOf((*MockRepository)(nil).FindChannelsMembershipsByMemberId), memberId)
}

// AddChannelMember mocks base method
func (m *MockRepository) AddChannelMember(mod *types.ChannelMember) (*types.ChannelMember, error) {
	ret := m.ctrl.Call(m, "AddChannelMember", mod)
	ret0, _ := ret[0].(*types.ChannelMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddChannelMember indicates an expected call of AddChannelMember
func (mr *MockRepositoryMockRecorder) AddChannelMember(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChannelMember", reflect.TypeOf((*MockRepository)(nil).AddChannelMember), mod)
}

// RemoveChannelMember mocks base method
func (m *MockRepository) RemoveChannelMember(channelID, userID uint64) error {
	ret := m.ctrl.Call(m, "RemoveChannelMember", channelID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveChannelMember indicates an expected call of RemoveChannelMember
func (mr *MockRepositoryMockRecorder) RemoveChannelMember(channelID, userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChannelMember", reflect.TypeOf((*MockRepository)(nil).RemoveChannelMember), channelID, userID)
}

// ArchiveChannelByID mocks base method
func (m *MockRepository) ArchiveChannelByID(id uint64) error {
	ret := m.ctrl.Call(m, "ArchiveChannelByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArchiveChannelByID indicates an expected call of ArchiveChannelByID
func (mr *MockRepositoryMockRecorder) ArchiveChannelByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveChannelByID", reflect.TypeOf((*MockRepository)(nil).ArchiveChannelByID), id)
}

// UnarchiveChannelByID mocks base method
func (m *MockRepository) UnarchiveChannelByID(id uint64) error {
	ret := m.ctrl.Call(m, "UnarchiveChannelByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnarchiveChannelByID indicates an expected call of UnarchiveChannelByID
func (mr *MockRepositoryMockRecorder) UnarchiveChannelByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveChannelByID", reflect.TypeOf((*MockRepository)(nil).UnarchiveChannelByID), id)
}

// DeleteChannelByID mocks base method
func (m *MockRepository) DeleteChannelByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteChannelByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChannelByID indicates an expected call of DeleteChannelByID
func (mr *MockRepositoryMockRecorder) DeleteChannelByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChannelByID", reflect.TypeOf((*MockRepository)(nil).DeleteChannelByID), id)
}

// FindMessageByID mocks base method
func (m *MockRepository) FindMessageByID(id uint64) (*types.Message, error) {
	ret := m.ctrl.Call(m, "FindMessageByID", id)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMessageByID indicates an expected call of FindMessageByID
func (mr *MockRepositoryMockRecorder) FindMessageByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMessageByID", reflect.TypeOf((*MockRepository)(nil).FindMessageByID), id)
}

// FindMessages mocks base method
func (m *MockRepository) FindMessages(filter *types.MessageFilter) ([]*types.Message, error) {
	ret := m.ctrl.Call(m, "FindMessages", filter)
	ret0, _ := ret[0].([]*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMessages indicates an expected call of FindMessages
func (mr *MockRepositoryMockRecorder) FindMessages(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMessages", reflect.TypeOf((*MockRepository)(nil).FindMessages), filter)
}

// CreateMessage mocks base method
func (m *MockRepository) CreateMessage(mod *types.Message) (*types.Message, error) {
	ret := m.ctrl.Call(m, "CreateMessage", mod)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage
func (mr *MockRepositoryMockRecorder) CreateMessage(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockRepository)(nil).CreateMessage), mod)
}

// UpdateMessage mocks base method
func (m *MockRepository) UpdateMessage(mod *types.Message) (*types.Message, error) {
	ret := m.ctrl.Call(m, "UpdateMessage", mod)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMessage indicates an expected call of UpdateMessage
func (mr *MockRepositoryMockRecorder) UpdateMessage(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockRepository)(nil).UpdateMessage), mod)
}

// DeleteMessageByID mocks base method
func (m *MockRepository) DeleteMessageByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteMessageByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageByID indicates an expected call of DeleteMessageByID
func (mr *MockRepositoryMockRecorder) DeleteMessageByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageByID", reflect.TypeOf((*MockRepository)(nil).DeleteMessageByID), id)
}

// FindOrganisationByID mocks base method
func (m *MockRepository) FindOrganisationByID(id uint64) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "FindOrganisationByID", id)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganisationByID indicates an expected call of FindOrganisationByID
func (mr *MockRepositoryMockRecorder) FindOrganisationByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganisationByID", reflect.TypeOf((*MockRepository)(nil).FindOrganisationByID), id)
}

// FindOrganisations mocks base method
func (m *MockRepository) FindOrganisations(filter *types.OrganisationFilter) ([]*types.Organisation, error) {
	ret := m.ctrl.Call(m, "FindOrganisations", filter)
	ret0, _ := ret[0].([]*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganisations indicates an expected call of FindOrganisations
func (mr *MockRepositoryMockRecorder) FindOrganisations(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganisations", reflect.TypeOf((*MockRepository)(nil).FindOrganisations), filter)
}

// CreateOrganisation mocks base method
func (m *MockRepository) CreateOrganisation(mod *types.Organisation) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "CreateOrganisation", mod)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganisation indicates an expected call of CreateOrganisation
func (mr *MockRepositoryMockRecorder) CreateOrganisation(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganisation", reflect.TypeOf((*MockRepository)(nil).CreateOrganisation), mod)
}

// UpdateOrganisation mocks base method
func (m *MockRepository) UpdateOrganisation(mod *types.Organisation) (*types.Organisation, error) {
	ret := m.ctrl.Call(m, "UpdateOrganisation", mod)
	ret0, _ := ret[0].(*types.Organisation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganisation indicates an expected call of UpdateOrganisation
func (mr *MockRepositoryMockRecorder) UpdateOrganisation(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganisation", reflect.TypeOf((*MockRepository)(nil).UpdateOrganisation), mod)
}

// ArchiveOrganisationByID mocks base method
func (m *MockRepository) ArchiveOrganisationByID(id uint64) error {
	ret := m.ctrl.Call(m, "ArchiveOrganisationByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArchiveOrganisationByID indicates an expected call of ArchiveOrganisationByID
func (mr *MockRepositoryMockRecorder) ArchiveOrganisationByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveOrganisationByID", reflect.TypeOf((*MockRepository)(nil).ArchiveOrganisationByID), id)
}

// UnarchiveOrganisationByID mocks base method
func (m *MockRepository) UnarchiveOrganisationByID(id uint64) error {
	ret := m.ctrl.Call(m, "UnarchiveOrganisationByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnarchiveOrganisationByID indicates an expected call of UnarchiveOrganisationByID
func (mr *MockRepositoryMockRecorder) UnarchiveOrganisationByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveOrganisationByID", reflect.TypeOf((*MockRepository)(nil).UnarchiveOrganisationByID), id)
}

// DeleteOrganisationByID mocks base method
func (m *MockRepository) DeleteOrganisationByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteOrganisationByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganisationByID indicates an expected call of DeleteOrganisationByID
func (mr *MockRepositoryMockRecorder) DeleteOrganisationByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganisationByID", reflect.TypeOf((*MockRepository)(nil).DeleteOrganisationByID), id)
}

// FindReactionByID mocks base method
func (m *MockRepository) FindReactionByID(id uint64) (*types.Reaction, error) {
	ret := m.ctrl.Call(m, "FindReactionByID", id)
	ret0, _ := ret[0].(*types.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReactionByID indicates an expected call of FindReactionByID
func (mr *MockRepositoryMockRecorder) FindReactionByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReactionByID", reflect.TypeOf((*MockRepository)(nil).FindReactionByID), id)
}

// FindReactionsByRange mocks base method
func (m *MockRepository) FindReactionsByRange(channelID, fromReactionID, toReactionID uint64) ([]*types.Reaction, error) {
	ret := m.ctrl.Call(m, "FindReactionsByRange", channelID, fromReactionID, toReactionID)
	ret0, _ := ret[0].([]*types.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReactionsByRange indicates an expected call of FindReactionsByRange
func (mr *MockRepositoryMockRecorder) FindReactionsByRange(channelID, fromReactionID, toReactionID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReactionsByRange", reflect.TypeOf((*MockRepository)(nil).FindReactionsByRange), channelID, fromReactionID, toReactionID)
}

// CreateReaction mocks base method
func (m *MockRepository) CreateReaction(mod *types.Reaction) (*types.Reaction, error) {
	ret := m.ctrl.Call(m, "CreateReaction", mod)
	ret0, _ := ret[0].(*types.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReaction indicates an expected call of CreateReaction
func (mr *MockRepositoryMockRecorder) CreateReaction(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReaction", reflect.TypeOf((*MockRepository)(nil).CreateReaction), mod)
}

// DeleteReactionByID mocks base method
func (m *MockRepository) DeleteReactionByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteReactionByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReactionByID indicates an expected call of DeleteReactionByID
func (mr *MockRepositoryMockRecorder) DeleteReactionByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReactionByID", reflect.TypeOf((*MockRepository)(nil).DeleteReactionByID), id)
}

// FindTeamByID mocks base method
func (m *MockRepository) FindTeamByID(id uint64) (*types.Team, error) {
	ret := m.ctrl.Call(m, "FindTeamByID", id)
	ret0, _ := ret[0].(*types.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTeamByID indicates an expected call of FindTeamByID
func (mr *MockRepositoryMockRecorder) FindTeamByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTeamByID", reflect.TypeOf((*MockRepository)(nil).FindTeamByID), id)
}

// FindTeams mocks base method
func (m *MockRepository) FindTeams(filter *types.TeamFilter) ([]*types.Team, error) {
	ret := m.ctrl.Call(m, "FindTeams", filter)
	ret0, _ := ret[0].([]*types.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTeams indicates an expected call of FindTeams
func (mr *MockRepositoryMockRecorder) FindTeams(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTeams", reflect.TypeOf((*MockRepository)(nil).FindTeams), filter)
}

// CreateTeam mocks base method
func (m *MockRepository) CreateTeam(mod *types.Team) (*types.Team, error) {
	ret := m.ctrl.Call(m, "CreateTeam", mod)
	ret0, _ := ret[0].(*types.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam
func (mr *MockRepositoryMockRecorder) CreateTeam(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockRepository)(nil).CreateTeam), mod)
}

// UpdateTeam mocks base method
func (m *MockRepository) UpdateTeam(mod *types.Team) (*types.Team, error) {
	ret := m.ctrl.Call(m, "UpdateTeam", mod)
	ret0, _ := ret[0].(*types.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeam indicates an expected call of UpdateTeam
func (mr *MockRepositoryMockRecorder) UpdateTeam(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeam", reflect.TypeOf((*MockRepository)(nil).UpdateTeam), mod)
}

// ArchiveTeamByID mocks base method
func (m *MockRepository) ArchiveTeamByID(id uint64) error {
	ret := m.ctrl.Call(m, "ArchiveTeamByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArchiveTeamByID indicates an expected call of ArchiveTeamByID
func (mr *MockRepositoryMockRecorder) ArchiveTeamByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveTeamByID", reflect.TypeOf((*MockRepository)(nil).ArchiveTeamByID), id)
}

// UnarchiveTeamByID mocks base method
func (m *MockRepository) UnarchiveTeamByID(id uint64) error {
	ret := m.ctrl.Call(m, "UnarchiveTeamByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnarchiveTeamByID indicates an expected call of UnarchiveTeamByID
func (mr *MockRepositoryMockRecorder) UnarchiveTeamByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveTeamByID", reflect.TypeOf((*MockRepository)(nil).UnarchiveTeamByID), id)
}

// DeleteTeamByID mocks base method
func (m *MockRepository) DeleteTeamByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteTeamByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeamByID indicates an expected call of DeleteTeamByID
func (mr *MockRepositoryMockRecorder) DeleteTeamByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeamByID", reflect.TypeOf((*MockRepository)(nil).DeleteTeamByID), id)
}

// MergeTeamByID mocks base method
func (m *MockRepository) MergeTeamByID(id, targetTeamID uint64) error {
	ret := m.ctrl.Call(m, "MergeTeamByID", id, targetTeamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeTeamByID indicates an expected call of MergeTeamByID
func (mr *MockRepositoryMockRecorder) MergeTeamByID(id, targetTeamID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeTeamByID", reflect.TypeOf((*MockRepository)(nil).MergeTeamByID), id, targetTeamID)
}

// MoveTeamByID mocks base method
func (m *MockRepository) MoveTeamByID(id, targetOrganisationID uint64) error {
	ret := m.ctrl.Call(m, "MoveTeamByID", id, targetOrganisationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveTeamByID indicates an expected call of MoveTeamByID
func (mr *MockRepositoryMockRecorder) MoveTeamByID(id, targetOrganisationID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveTeamByID", reflect.TypeOf((*MockRepository)(nil).MoveTeamByID), id, targetOrganisationID)
}

// FindUserByUsername mocks base method
func (m *MockRepository) FindUserByUsername(username string) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindUserByUsername", username)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsername indicates an expected call of FindUserByUsername
func (mr *MockRepositoryMockRecorder) FindUserByUsername(username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsername", reflect.TypeOf((*MockRepository)(nil).FindUserByUsername), username)
}

// FindUserByID mocks base method
func (m *MockRepository) FindUserByID(id uint64) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindUserByID", id)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID
func (mr *MockRepositoryMockRecorder) FindUserByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockRepository)(nil).FindUserByID), id)
}

// FindUsers mocks base method
func (m *MockRepository) FindUsers(filter *types.UserFilter) ([]*types.User, error) {
	ret := m.ctrl.Call(m, "FindUsers", filter)
	ret0, _ := ret[0].([]*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsers indicates an expected call of FindUsers
func (mr *MockRepositoryMockRecorder) FindUsers(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsers", reflect.TypeOf((*MockRepository)(nil).FindUsers), filter)
}

// CreateUser mocks base method
func (m *MockRepository) CreateUser(mod *types.User) (*types.User, error) {
	ret := m.ctrl.Call(m, "CreateUser", mod)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockRepositoryMockRecorder) CreateUser(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), mod)
}

// UpdateUser mocks base method
func (m *MockRepository) UpdateUser(mod *types.User) (*types.User, error) {
	ret := m.ctrl.Call(m, "UpdateUser", mod)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockRepositoryMockRecorder) UpdateUser(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), mod)
}

// SuspendUserByID mocks base method
func (m *MockRepository) SuspendUserByID(id uint64) error {
	ret := m.ctrl.Call(m, "SuspendUserByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SuspendUserByID indicates an expected call of SuspendUserByID
func (mr *MockRepositoryMockRecorder) SuspendUserByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendUserByID", reflect.TypeOf((*MockRepository)(nil).SuspendUserByID), id)
}

// UnsuspendUserByID mocks base method
func (m *MockRepository) UnsuspendUserByID(id uint64) error {
	ret := m.ctrl.Call(m, "UnsuspendUserByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsuspendUserByID indicates an expected call of UnsuspendUserByID
func (mr *MockRepositoryMockRecorder) UnsuspendUserByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsuspendUserByID", reflect.TypeOf((*MockRepository)(nil).UnsuspendUserByID), id)
}

// DeleteUserByID mocks base method
func (m *MockRepository) DeleteUserByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteUserByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserByID indicates an expected call of DeleteUserByID
func (mr *MockRepositoryMockRecorder) DeleteUserByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockRepository)(nil).DeleteUserByID), id)
}

// EventQueuePull mocks base method
func (m *MockRepository) EventQueuePull(origin uint64) ([]*types.EventQueueItem, error) {
	ret := m.ctrl.Call(m, "EventQueuePull", origin)
	ret0, _ := ret[0].([]*types.EventQueueItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventQueuePull indicates an expected call of EventQueuePull
func (mr *MockRepositoryMockRecorder) EventQueuePull(origin interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventQueuePull", reflect.TypeOf((*MockRepository)(nil).EventQueuePull), origin)
}

// EventQueuePush mocks base method
func (m *MockRepository) EventQueuePush(eqi *types.EventQueueItem) error {
	ret := m.ctrl.Call(m, "EventQueuePush", eqi)
	ret0, _ := ret[0].(error)
	return ret0
}

// EventQueuePush indicates an expected call of EventQueuePush
func (mr *MockRepositoryMockRecorder) EventQueuePush(eqi interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventQueuePush", reflect.TypeOf((*MockRepository)(nil).EventQueuePush), eqi)
}

// EventQueueSync mocks base method
func (m *MockRepository) EventQueueSync(origin, id uint64) error {
	ret := m.ctrl.Call(m, "EventQueueSync", origin, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// EventQueueSync indicates an expected call of EventQueueSync
func (mr *MockRepositoryMockRecorder) EventQueueSync(origin, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventQueueSync", reflect.TypeOf((*MockRepository)(nil).EventQueueSync), origin, id)
}
