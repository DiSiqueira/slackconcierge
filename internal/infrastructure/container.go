package infrastructure

import (
	"errors"
	"github.com/disiqueira/SlackConcierge/internal/application"
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"github.com/disiqueira/SlackConcierge/internal/handler"
	"github.com/disiqueira/SlackConcierge/internal/infrastructure/console"
	"github.com/disiqueira/SlackConcierge/internal/infrastructure/thirdparty"
	"github.com/disiqueira/SlackConcierge/internal/infrastructure/file"
)

// Container is a global structure for holding references to application resources
type Container struct {
	configurationService application.ConfigurationService
	handlerService       application.HandlerService
	logService           application.LogService
	slackService         application.SlackService

	logRepository           domain.LogRepository
	messageRepository       domain.MessageRepository
	configurationRepository domain.ConfigurationRepository
}

func (c *Container) ConfigurationService() application.ConfigurationService {
	if c.configurationService == nil {
		c.configurationService = application.NewConfigurationService(c.ConfigurationRepository())
	}

	return c.configurationService
}

func (c *Container) ConfigurationRepository() domain.ConfigurationRepository {
	if c.configurationRepository == nil {
		c.configurationRepository = file.NewConfigurationRepository(c.LogService())
	}

	return c.configurationRepository
}

func (c *Container) LogService() application.LogService {
	if c.logService == nil {
		c.logService = application.NewLogService(c.LogRepository())
	}

	return c.logService
}

func (c *Container) LogRepository() domain.LogRepository {
	if c.logRepository == nil {
		c.logRepository = console.NewLogRepository()
	}

	return c.logRepository
}

func (c *Container) SlackService() application.SlackService {
	if c.slackService == nil {
		c.slackService = application.NewSlackService(
			c.MessageRepository(),
			c.ConfigurationService().GetIntDefault("slack_buffer", 64),
			)
	}

	return c.slackService
}

func (c *Container) MessageRepository() domain.MessageRepository {
	if c.messageRepository == nil {
		bufferSize := c.ConfigurationService().GetIntDefault("slack_buffer", 64)

		slackToken := c.ConfigurationService().Get("slack_token")
		if slackToken == "" {
			c.LogService().Fatal(errors.New("slack token not found"))
		}

		slackRepository, err := thirdparty.NewSlackRepository(slackToken, bufferSize)
		if err != nil {
			c.LogService().Fatal(err)
		}
		c.messageRepository = slackRepository
	}

	return c.messageRepository
}

func (c *Container) HandlerService() application.HandlerService {
	if c.handlerService == nil {
		c.handlerService = application.NewHandlerService([]handler.Handler{
				handler.Log{},
				handler.Program{},
			})

	}
	return c.handlerService
}
