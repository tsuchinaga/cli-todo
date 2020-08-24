package repository

import "time"

type ClockRepository interface {
	Now() time.Time
}
