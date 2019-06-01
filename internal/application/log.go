//go:generate mockgen -source log.go -destination log_mock.go -package application
package application

import "github.com/disiqueira/SlackConcierge/internal/domain"

type (
	logService struct {
		logRepository domain.LogRepository
	}

	LogService interface {
		Error(err error)
		Notice(err error)
		Fatal(err error)
		Errors(errs []error)
	}
)

func NewLogService(repository domain.LogRepository) LogService {
	return &logService{
		logRepository: repository,
	}
}

func (l *logService) Error(err error) {
	if err == nil {
		return
	}
	l.logRepository.Error(&domain.Log{
		Message: err.Error(),
	})
}

func (l *logService) Errors(errs []error) {
	if len(errs) == 0 {
		return
	}

	l.Error(errs[0])
	l.Errors(errs[1:])
}

func (l *logService) Fatal(err error) {
	if err == nil {
		return
	}
	l.logRepository.Fatal(&domain.Log{
		Message: err.Error(),
	})
}

func (l *logService) Notice(err error) {
	l.logRepository.Notice(&domain.Log{
		Message: err.Error(),
	})
}
