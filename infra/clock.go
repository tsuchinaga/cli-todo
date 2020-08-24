package infra

import (
	"time"

	"gitlab.com/tsuchinaga/cli-todo/app/repository"
)

func NewClock() repository.ClockRepository {
	return &clock{}
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
}
