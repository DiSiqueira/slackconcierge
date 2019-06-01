package file

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/disiqueira/SlackConcierge/internal/application"
)

type (
	ConfigurationRepository struct {
		LogService application.LogService
	}
)

func NewConfigurationRepository(logService application.LogService) *ConfigurationRepository {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/slack_concierge/")
	viper.AddConfigPath("$HOME/.slack_concierge")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logService.Notice(err)
	}

	return &ConfigurationRepository{
		LogService: logService,
	}
}

func (c *ConfigurationRepository) Get(name string) string {
	return viper.GetString(strings.ToUpper(name))
}

func (c *ConfigurationRepository) Set(name string, value interface{}) {
 	viper.Set(strings.ToUpper(name), value)
}

func (c *ConfigurationRepository) GetInt(name string) int {
	return viper.GetInt(strings.ToUpper(name))
}
