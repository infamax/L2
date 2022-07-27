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

func TestGetEventsForDay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	mockRes := []models.Event{
		{
			UserID:       1,
			EventID:      1,
			Title:        "important",
			DateCreated:  "2022-07-27",
			DateFinished: "2022-07-28",
			Description:  "Passed exam L2",
			Done:         false,
		},
		{
			UserID:       2,
			EventID:      2,
			Title:        "not important",
			DateCreated:  "2022-07-27",
			DateFinished: "2022-07-28",
			Description:  "Make cleaning",
			Done:         false,
		},
	}
	userID := 1
	date := "2022-07-27"
	count := 1
	ctx := context.Background()
	repo.EXPECT().GetEvents(ctx, userID, date, count).Return(mockRes, nil).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	res, err := srv.GetEventsForDay(ctx, userID, date)
	assert.Nil(t, err)
	assert.Equal(t, mockRes, res)
}

func TestGetInvalidEventsForDay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := mockRepo.NewMockRepository(ctl)
	userID := 1
	date := "2022-07-27"
	count := 1
	ctx := context.Background()
	mockErr := errors.New("no such user in repo")
	repo.EXPECT().GetEvents(ctx, userID, date, count).Return([]models.Event{}, mockErr).Times(1)
	srv, err := service.New(repo)
	assert.Nil(t, err)
	res, err := srv.GetEventsForDay(ctx, userID, date)
	assert.EqualError(t, err, mockErr.Error())
	assert.Equal(t, []models.Event{}, res)
}
