package scheduler_status

import (
	"github.com/google/uuid"
	"time"
)

const (
	StateReady    ScheduleState = "ready"
	StateNotReady ScheduleState = "not_ready"
)

type ScheduleState string

type SchedulerStatus struct {
	ID        uuid.UUID     `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	State     ScheduleState `json:"state" db:"state"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}
