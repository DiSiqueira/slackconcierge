package application_test

import (
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"github.com/disiqueira/SlackConcierge/internal/handler"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/disiqueira/SlackConcierge/internal/application"
)

func TestHandlerService_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockHandler := handler.NewMockHandler(ctrl)
	handlerList := []handler.Handler{
		mockHandler,
	}

	handlerService := application.NewHandlerService(handlerList)

	msg := &domain.Message{
		Text: "test Message",
		Channel: "testChannel",
	}

	mockHandler.EXPECT().Execute(msg).Return([]*domain.Message{msg}, nil)

	gotResponse, gotErrors := handlerService.Handle(msg)
	if len(gotErrors) > 0 {
		t.Errorf("error returned %s", gotErrors[0])
	}

	if total := len(gotResponse); total != 1 {
		t.Errorf("expecting 1 message got %d", total)
	}
}
