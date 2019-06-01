//go:generate mockgen -source configuration.go -destination configuration_mock.go -package domain
package domain

type (
	Configuration struct {
		Value string
	}

	ConfigurationRepository interface {
		Get(name string) string
		Set(name string, value interface{})
		GetInt(name string) int
	}
)
