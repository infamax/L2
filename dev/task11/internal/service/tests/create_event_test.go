package tests

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/infamax/l2/task11/internal/models"
	"github.com/infamax/l2/task11/internal/service"
	mockRepo "github.com/infamax/l2/task11/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEventValid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	ctx := context.Background()
	mockResp := 1
	event := &models.Event{
		UserID:       1,
		Title:        "important",
		DateCreated:  "2022-07-27",
		DateFinished: "2022-07-28",
		Description:  "Pass exam L2",
		Done:         false,
	}
	repo.EXPECT().CreateEvent(ctx, event).Return(mockResp, nil).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	expectedRes := 1
	res, err := srv.CreateEvent(ctx, event)
	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
}

func TestCreateInvalidData(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	event := &models.Event{
		UserID:       1,
		Title:        "important",
		DateCreated:  "2022-07-27",
		DateFinished: "2022-07-28",
		Description:  "Pass exam L2",
		Done:         false,
	}
	mockErrRes := errors.New("this user already exists in db")
	mockReturningRes := 0
	ctx := context.Background()
	repo.EXPECT().CreateEvent(ctx, event).Return(mockReturningRes, mockErrRes).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	res, err := srv.CreateEvent(ctx, event)
	expectedRes := 0
	assert.EqualError(t, err, mockErrRes.Error())
	assert.Equal(t, expectedRes, res)
}
