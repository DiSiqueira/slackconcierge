package application_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/disiqueira/SlackConcierge/internal/application"
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"github.com/disiqueira/SlackConcierge/internal/infrastructure/console"
)

func TestLogService_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	err := errors.New("test error")

	mockLogRepository := console.NewMockLogRepository(ctrl)
	mockLogRepository.EXPECT().Error(&domain.Log{
		Message: err.Error(),
	})

	logService := application.NewLogService(mockLogRepository)

	logService.Error(err)
}

func TestLogService_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)

	err := errors.New("test error")

	mockLogRepository := console.NewMockLogRepository(ctrl)
	mockLogRepository.EXPECT().Error(&domain.Log{
		Message: err.Error(),
	}).Times(2)

	logService := application.NewLogService(mockLogRepository)

	logService.Errors([]error{err, err})
}

func TestLogService_Fatal(t *testing.T) {
	ctrl := gomock.NewController(t)

	err := errors.New("test error")

	mockLogRepository := console.NewMockLogRepository(ctrl)
	mockLogRepository.EXPECT().Fatal(&domain.Log{
		Message: err.Error(),
	})

	logService := application.NewLogService(mockLogRepository)

	logService.Fatal(err)
}

func TestLogService_Notice(t *testing.T) {
	ctrl := gomock.NewController(t)

	err := errors.New("test error")

	mockLogRepository := console.NewMockLogRepository(ctrl)
	mockLogRepository.EXPECT().Notice(&domain.Log{
		Message: err.Error(),
	})

	logService := application.NewLogService(mockLogRepository)

	logService.Notice(err)
}
