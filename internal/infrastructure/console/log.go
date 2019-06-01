package console

import (
	"fmt"
	"github.com/disiqueira/SlackConcierge/internal/domain"
	"log"
)

type (
	LogRepository struct {

	}
)

func NewLogRepository() *LogRepository {
	return &LogRepository{}
}

func (l *LogRepository) Error(err *domain.Log) {
	fmt.Println("Error: " + err.Message)
}

func (l *LogRepository) Fatal(err *domain.Log) {
	log.Fatalf("Fatal: %s", err.Message)
}

func (l *LogRepository) Notice(err *domain.Log) {
	fmt.Println("Notice: " + err.Message)
}
