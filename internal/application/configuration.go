//go:generate mockgen -source configuration.go -destination configuration_mock.go -package application
package application

import "github.com/disiqueira/SlackConcierge/internal/domain"

type (
	configurationService struct {
		configurationRepository domain.ConfigurationRepository
	}

	ConfigurationService interface {
		GetIntDefault(name string, fallback int) int
		GetInt(name string) int
		GetDefault(name, fallback string) string
		Get(name string) string
		Set(name string, value interface{})
	}
)

func NewConfigurationService(repository domain.ConfigurationRepository) ConfigurationService {
	return &configurationService{
		configurationRepository: repository,
	}
}

func (c *configurationService) Get(name string) string {
	return c.configurationRepository.Get(name)
}

func (c *configurationService) GetDefault(name, fallback string) string {
	if val := c.Get(name); val != "" {
		return val
	}
	c.Set(name, fallback)
	return fallback
}

func (c *configurationService) GetInt(name string) int {
	return c.configurationRepository.GetInt(name)
}

func (c *configurationService) GetIntDefault(name string, fallback int) int {
	if val := c.GetInt(name); val != 0 {
		return val
	}
	c.Set(name, fallback)
	return fallback
}
func (c *configurationService) Set(name string, value interface{}) {
	c.configurationRepository.Set(name, value)
}
