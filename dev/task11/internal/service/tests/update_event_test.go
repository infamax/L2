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

func TestUpdateEventValid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	event := &models.Event{
		UserID:       1,
		Title:        "important",
		DateCreated:  "2022-07-27",
		DateFinished: "2022-07-28",
		Description:  "Pass exam L2",
		Done:         true,
	}
	ctx := context.Background()
	repo.EXPECT().UpdateEvent(ctx, event).Return(nil).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	err = srv.UpdateEvent(ctx, event)
	assert.Nil(t, err)
}

func TestUpdateEventInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	event := &models.Event{
		UserID:       1,
		Title:        "important",
		DateCreated:  "2022-07-27",
		DateFinished: "2022-07-28",
		Description:  "Pass exam L2",
		Done:         true,
	}
	ctx := context.Background()
	mockErr := errors.New("no such event with id in repo")
	repo.EXPECT().UpdateEvent(ctx, event).Return(mockErr).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	err = srv.UpdateEvent(ctx, event)
	assert.EqualError(t, err, mockErr.Error())
}
