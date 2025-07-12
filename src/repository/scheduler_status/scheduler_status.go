package scheduler_status

import (
	"context"
	"github.com/jmoiron/sqlx"
	"unit_of_work__go/domain/scheduler_status"
	"unit_of_work__go/infra/db"
)

type RepositorySchedulerStatus struct {
	db *sqlx.DB
}

func NewRepositorySchedulerStatus(dbClient *sqlx.DB) *RepositorySchedulerStatus {
	return &RepositorySchedulerStatus{
		db: dbClient,
	}
}

func (ss *RepositorySchedulerStatus) Read(ctx context.Context, q db.Querier) ([]*scheduler_status.SchedulerStatus, error) {
	var schedulerStatusResult []scheduler_status.SchedulerStatus
	query := `SELECT id, name, state, created_at, updated_at FROM scheduler_status;`
	err := q.SelectContext(ctx, &schedulerStatusResult, query)
	if err != nil {
		return nil, err
	}

	result := make([]*scheduler_status.SchedulerStatus, len(schedulerStatusResult))
	if len(schedulerStatusResult) > 0 {
		for i := range schedulerStatusResult {
			result[i] = &schedulerStatusResult[i]
		}
	}
	return result, nil
}

func (ss *RepositorySchedulerStatus) Load(ctx context.Context, q db.Querier, status *scheduler_status.SchedulerStatus) error {
	sqlQuery := `
		INSERT INTO scheduler_status (id, name, state, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err := q.ExecContext(
		ctx,
		sqlQuery,
		status.ID,
		status.Name,
		status.State,
		status.CreatedAt,
		status.UpdatedAt,
	)
	return err
}

func (ss *RepositorySchedulerStatus) DeleteAll(ctx context.Context, q db.Querier) error {
	sqlQuery := `DELETE FROM scheduler_status;`
	_, err := q.ExecContext(ctx, sqlQuery)
	return err
}
