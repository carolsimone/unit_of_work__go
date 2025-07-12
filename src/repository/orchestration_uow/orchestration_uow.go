package orchestration_uow

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"unit_of_work__go/infra/db"
	"unit_of_work__go/repository/k8s_deployment_status"
	"unit_of_work__go/repository/scheduler_status"
)

type OrchestrationUnitOfWork struct {
	db *sqlx.DB
}

func NewOrchestrationUnitOfWork(db *sqlx.DB) (*OrchestrationUnitOfWork, error) {
	return &OrchestrationUnitOfWork{db: db}, nil
}

func (uow *OrchestrationUnitOfWork) RunInTx(ctx context.Context, fn func(ctx context.Context, tx db.Querier) error) error {
	// Here you create the transaction which is then passed as db.Querier in the function at the very end
	tx, err := uow.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				// Log the rollback error if needed
				log.Printf("Failed to rollback transaction: %v", err)
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				// Log the rollback error if needed
				log.Printf("Failed to rollback transaction: %v", err)
			}
		} else {
			err = tx.Commit()
		}
	}()
	return fn(ctx, tx)
}

func (uow *OrchestrationUnitOfWork) K8sDeploymentRepo() *k8s_deployment_status.RepositoryK8sDeploymentStatus {
	return k8s_deployment_status.NewRepositoryDeploymentStatus(uow.db)
}

func (uow *OrchestrationUnitOfWork) SchedulerStatusRepo() *scheduler_status.RepositorySchedulerStatus {
	return scheduler_status.NewRepositorySchedulerStatus(uow.db)
}

// DB returns the underlying database connection for direct queries if needed.
func (uow *OrchestrationUnitOfWork) DB() db.Querier {
	return uow.db
}
