package k8s_deployment_status

import (
	"context"
	"github.com/jmoiron/sqlx"
	"unit_of_work__go/domain/k8s_deployment_status"
	"unit_of_work__go/infra/db"
)

type RepositoryK8sDeploymentStatus struct {
	db *sqlx.DB
}

func NewRepositoryDeploymentStatus(dbClient *sqlx.DB) *RepositoryK8sDeploymentStatus {
	return &RepositoryK8sDeploymentStatus{
		db: dbClient,
	}
}

func (kds *RepositoryK8sDeploymentStatus) Read(ctx context.Context, q db.Querier) ([]*k8s_deployment_status.K8sDeploymentStatus, error) {
	var k8sDeploymentStatusResult []k8s_deployment_status.K8sDeploymentStatus
	query := `select id, scheduler_id, state, created_at, updated_at from k8s_deployment_status;`
	err := q.SelectContext(ctx, &k8sDeploymentStatusResult, query)
	if err != nil {
		return nil, err
	}
	result := make([]*k8s_deployment_status.K8sDeploymentStatus, len(k8sDeploymentStatusResult))
	if len(k8sDeploymentStatusResult) > 0 {
		for i := range k8sDeploymentStatusResult {
			result[i] = &k8sDeploymentStatusResult[i]
		}
	}
	return result, nil
}

func (kds *RepositoryK8sDeploymentStatus) Load(ctx context.Context, q db.Querier, k8sDeploymentStatus *k8s_deployment_status.K8sDeploymentStatus) error {
	// Extract transaction from context if needed
	query := `
		insert into k8s_deployment_status (id, scheduler_id, state, created_at, updated_at) 
		values ($1, $2, $3, $4, $5);
	`
	_, err := q.ExecContext(
		ctx,
		query,
		k8sDeploymentStatus.ID,
		k8sDeploymentStatus.SchedulerID,
		k8sDeploymentStatus.State,
		k8sDeploymentStatus.CreatedAt,
		k8sDeploymentStatus.UpdatedAt,
	)
	return err
}

func (kds *RepositoryK8sDeploymentStatus) DeleteAll(ctx context.Context, q db.Querier) error {
	query := `delete from k8s_deployment_status;`
	_, err := q.ExecContext(ctx, query)
	return err
}
