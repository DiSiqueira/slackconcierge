package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/disiqueira/SlackConcierge/internal/application"
	"github.com/disiqueira/SlackConcierge/internal/domain"
)

func TestConfigurationService_Get(t *testing.T) {
	key := "key1"
	val := "value1"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := domain.NewMockConfigurationRepository(ctrl)

	mockRepository.EXPECT().Get(key).Return(val)

	config := application.NewConfigurationService(mockRepository)
	got := config.Get(key)

	if got != val {
		t.Errorf("expecting %s got %s", val, got)
	}
}

func TestConfigurationService_GetDefault(t *testing.T) {
	key := "key1"
	val := "value1"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := domain.NewMockConfigurationRepository(ctrl)

	mockRepository.EXPECT().Get(key).Return("")
	mockRepository.EXPECT().Set(key, val)

	config := application.NewConfigurationService(mockRepository)
	got := config.GetDefault(key, val)

	if got != val {
		t.Errorf("expecting %s got %s", val, got)
	}
}

func TestConfigurationService_GetInt(t *testing.T) {
	key := "key1"
	val := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := domain.NewMockConfigurationRepository(ctrl)

	mockRepository.EXPECT().GetInt(key).Return(val)

	config := application.NewConfigurationService(mockRepository)
	got := config.GetInt(key)

	if got != val {
		t.Errorf("expecting %d got %d", val, got)
	}
}

func TestConfigurationService_GetIntDefault(t *testing.T) {
	key := "key1"
	val := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := domain.NewMockConfigurationRepository(ctrl)

	mockRepository.EXPECT().GetInt(key).Return(0)
	mockRepository.EXPECT().Set(key, val)

	config := application.NewConfigurationService(mockRepository)
	got := config.GetIntDefault(key,val)

	if got != val {
		t.Errorf("expecting %d got %d", val, got)
	}
}

func TestConfigurationService_Set(t *testing.T) {
	key := "key1"
	val := 1

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := domain.NewMockConfigurationRepository(ctrl)

	mockRepository.EXPECT().Set(key, val)

	config := application.NewConfigurationService(mockRepository)
	config.Set(key,val)
}
