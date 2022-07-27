package tests

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/infamax/l2/task11/internal/service"
	mockRepo "github.com/infamax/l2/task11/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteEventValid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	eventID := 1
	ctx := context.Background()
	repo.EXPECT().DeleteEvent(ctx, eventID).Return(nil).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	err = srv.DeleteEvent(ctx, eventID)
	assert.Nil(t, err)
}

func TestDeleteNotValidEvent(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	eventID := 1
	ctx := context.Background()
	mockErr := errors.New("no event with such id in database")
	repo.EXPECT().DeleteEvent(ctx, eventID).Return(mockErr).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	err = srv.DeleteEvent(ctx, eventID)
	assert.EqualError(t, err, mockErr.Error())
}
