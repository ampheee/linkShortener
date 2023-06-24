package mockRepo

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"grpcService/pkg/utilities"
	"testing"
)

func TestMockRepo_InsertLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	originalLink := "https://www.example.com"
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(originalLink))
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(nil)
	err := mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.NoError(t, err, "unexpected error")
}

func TestInsertLink_EmptyLink(t *testing.T) {
	var Error = errors.New("ErrEmptyLink")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	mockRepo.EXPECT().InsertLink(gomock.Any(), "", "").Return(Error)
	err := mockRepo.InsertLink(context.Background(), "", "")
	require.Error(t, err, "expected error")
}

func TestInsertLink_LinkWithoutProtocol(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	originalLink := "www.example.com"
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(originalLink))
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(nil)
	err := mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.NoError(t, err, "unexpected error")
}

func TestInsertLink_ExistingLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	originalLink := "https://www.example.com"
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(originalLink))
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(nil)
	err := mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.NoError(t, err, "unexpected error")
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(errors.New("link already exists"))
	err = mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.Error(t, err, "expected error")
}

func TestInsertLink_InvalidLinkFormat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	originalLink := "example"
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(originalLink))
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(errors.New("invalid link format"))
	err := mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.Error(t, err, "expected error")
}

func TestInsertLink_LongLinkPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	originalLink := "https://www.example.com/this/is/a/very/long/path/and/should/be/handled/correctly"
	abbreviatedLink := utilities.EncodeBase63(utilities.HashLink(originalLink))
	mockRepo.EXPECT().InsertLink(gomock.Any(), abbreviatedLink, originalLink).Return(nil)
	err := mockRepo.InsertLink(context.Background(), abbreviatedLink, originalLink)
	require.NoError(t, err, "unexpected error")
}

func TestMockRepo_SelectLinkSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	abbreviatedLink := "abc"
	originalLink := "https://www.example.com"
	mockRepo.EXPECT().SelectLink(gomock.Any(), abbreviatedLink).Return(originalLink, nil)
	link, err := mockRepo.SelectLink(context.Background(), abbreviatedLink)
	require.NoError(t, err, "unexpected error")
	require.Equal(t, originalLink, link, "wrong link")
}

func TestMockRepo_SelectLink_InvalidAbbreviatedLink(t *testing.T) {
	var ErrLinkNotFound = errors.New("ErrInvalidAbbreviatedLink")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	abbreviatedLink := "invalid"
	mockRepo.EXPECT().SelectLink(gomock.Any(), abbreviatedLink).Return("", ErrLinkNotFound)
	_, err := mockRepo.SelectLink(context.Background(), abbreviatedLink)
	require.Error(t, err, "expected error")
	require.Equal(t, ErrLinkNotFound, err, "wrong error type")
}

func TestMockRepo_SelectLink_EmptyAbbreviatedLink(t *testing.T) {
	var ErrInvalidAbbreviatedLink = errors.New("ErrInvalidAbbreviatedLink")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	abbreviatedLink := ""
	mockRepo.EXPECT().SelectLink(gomock.Any(), abbreviatedLink).Return("", ErrInvalidAbbreviatedLink)
	_, err := mockRepo.SelectLink(context.Background(), abbreviatedLink)
	require.Error(t, err, "expected error")
	require.Equal(t, ErrInvalidAbbreviatedLink, err, "wrong error type")
}

func TestMockRepo_SelectLink_EmptyResult(t *testing.T) {
	var ErrLinkNotFound = errors.New("ErrLinkNotFound")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	abbreviatedLink := "abc"
	mockRepo.EXPECT().SelectLink(gomock.Any(), abbreviatedLink).Return("", ErrLinkNotFound)
	_, err := mockRepo.SelectLink(context.Background(), abbreviatedLink)
	require.Error(t, err, "expected error")
	require.Equal(t, ErrLinkNotFound, err, "wrong error type")
}

func TestMockRepo_SelectLink_DatabaseError(t *testing.T) {
	var ErrDatabase = errors.New("ErrDatabase")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepo(ctrl)
	abbreviatedLink := "abc"
	mockRepo.EXPECT().SelectLink(gomock.Any(), abbreviatedLink).Return("", ErrDatabase)
	_, err := mockRepo.SelectLink(context.Background(), abbreviatedLink)
	require.Error(t, err, "expected error")
	require.Equal(t, ErrDatabase, err, "wrong error type")
}
