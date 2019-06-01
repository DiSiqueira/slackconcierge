//go:generate mockgen -source log.go -destination log_mock.go -package domain
package domain

type (
	Log struct {
		Message string
	}

	LogRepository interface {
		Error(log *Log)
		Fatal(log *Log)
		Notice(log *Log)
	}
)

