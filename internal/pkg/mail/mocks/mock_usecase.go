// Code generated by MockGen. DO NOT EDIT.
// Source: liokor_mail/internal/pkg/mail (interfaces: MailUseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	mail "liokor_mail/internal/pkg/mail"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockMailUseCase is a mock of MailUseCase interface.
type MockMailUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMailUseCaseMockRecorder
}

// MockMailUseCaseMockRecorder is the mock recorder for MockMailUseCase.
type MockMailUseCaseMockRecorder struct {
	mock *MockMailUseCase
}

// NewMockMailUseCase creates a new mock instance.
func NewMockMailUseCase(ctrl *gomock.Controller) *MockMailUseCase {
	mock := &MockMailUseCase{ctrl: ctrl}
	mock.recorder = &MockMailUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailUseCase) EXPECT() *MockMailUseCaseMockRecorder {
	return m.recorder
}

// CreateDialogue mocks base method.
func (m *MockMailUseCase) CreateDialogue(arg0, arg1 string) (mail.Dialogue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDialogue", arg0, arg1)
	ret0, _ := ret[0].(mail.Dialogue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDialogue indicates an expected call of CreateDialogue.
func (mr *MockMailUseCaseMockRecorder) CreateDialogue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDialogue", reflect.TypeOf((*MockMailUseCase)(nil).CreateDialogue), arg0, arg1)
}

// CreateFolder mocks base method.
func (m *MockMailUseCase) CreateFolder(arg0 int, arg1 string) (mail.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", arg0, arg1)
	ret0, _ := ret[0].(mail.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFolder indicates an expected call of CreateFolder.
func (mr *MockMailUseCaseMockRecorder) CreateFolder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockMailUseCase)(nil).CreateFolder), arg0, arg1)
}

// DeleteDialogue mocks base method.
func (m *MockMailUseCase) DeleteDialogue(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDialogue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDialogue indicates an expected call of DeleteDialogue.
func (mr *MockMailUseCaseMockRecorder) DeleteDialogue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDialogue", reflect.TypeOf((*MockMailUseCase)(nil).DeleteDialogue), arg0, arg1)
}

// DeleteFolder mocks base method.
func (m *MockMailUseCase) DeleteFolder(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockMailUseCaseMockRecorder) DeleteFolder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockMailUseCase)(nil).DeleteFolder), arg0, arg1, arg2)
}

// DeleteMails mocks base method.
func (m *MockMailUseCase) DeleteMails(arg0 string, arg1 []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMails", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMails indicates an expected call of DeleteMails.
func (mr *MockMailUseCaseMockRecorder) DeleteMails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMails", reflect.TypeOf((*MockMailUseCase)(nil).DeleteMails), arg0, arg1)
}

// GetDialogues mocks base method.
func (m *MockMailUseCase) GetDialogues(arg0 string, arg1 int, arg2 string, arg3 int, arg4 time.Time) ([]mail.Dialogue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogues", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]mail.Dialogue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogues indicates an expected call of GetDialogues.
func (mr *MockMailUseCaseMockRecorder) GetDialogues(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogues", reflect.TypeOf((*MockMailUseCase)(nil).GetDialogues), arg0, arg1, arg2, arg3, arg4)
}

// GetEmails mocks base method.
func (m *MockMailUseCase) GetEmails(arg0, arg1 string, arg2, arg3 int) ([]mail.DialogueEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmails", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]mail.DialogueEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmails indicates an expected call of GetEmails.
func (mr *MockMailUseCaseMockRecorder) GetEmails(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmails", reflect.TypeOf((*MockMailUseCase)(nil).GetEmails), arg0, arg1, arg2, arg3)
}

// GetFolders mocks base method.
func (m *MockMailUseCase) GetFolders(arg0 int) ([]mail.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFolders", arg0)
	ret0, _ := ret[0].([]mail.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFolders indicates an expected call of GetFolders.
func (mr *MockMailUseCaseMockRecorder) GetFolders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFolders", reflect.TypeOf((*MockMailUseCase)(nil).GetFolders), arg0)
}

// SendEmail mocks base method.
func (m *MockMailUseCase) SendEmail(arg0 mail.Mail) (mail.Mail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", arg0)
	ret0, _ := ret[0].(mail.Mail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockMailUseCaseMockRecorder) SendEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockMailUseCase)(nil).SendEmail), arg0)
}

// UpdateFolderName mocks base method.
func (m *MockMailUseCase) UpdateFolderName(arg0, arg1 int, arg2 string) (mail.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFolderName", arg0, arg1, arg2)
	ret0, _ := ret[0].(mail.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFolderName indicates an expected call of UpdateFolderName.
func (mr *MockMailUseCaseMockRecorder) UpdateFolderName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFolderName", reflect.TypeOf((*MockMailUseCase)(nil).UpdateFolderName), arg0, arg1, arg2)
}

// UpdateFolderPutDialogue mocks base method.
func (m *MockMailUseCase) UpdateFolderPutDialogue(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFolderPutDialogue", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFolderPutDialogue indicates an expected call of UpdateFolderPutDialogue.
func (mr *MockMailUseCaseMockRecorder) UpdateFolderPutDialogue(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFolderPutDialogue", reflect.TypeOf((*MockMailUseCase)(nil).UpdateFolderPutDialogue), arg0, arg1, arg2)
}
