package k8s_deployment_status

import (
	"github.com/google/uuid"
	"time"
)

const (
	StatePending   K8sDeploymentState = "pending"
	StateRunning   K8sDeploymentState = "running"
	StateCompleted K8sDeploymentState = "completed"
	StateFailed    K8sDeploymentState = "failed"
)

type K8sDeploymentState string

// K8sDeploymentStatus This works both as a table to read from DB and as a domain model
type K8sDeploymentStatus struct {
	ID          uuid.UUID          `json:"id" db:"id"`
	SchedulerID uuid.UUID          `json:"scheduler_id" db:"scheduler_id"`
	State       K8sDeploymentState `json:"state" db:"state"`
	CreatedAt   time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" db:"updated_at"`
}
