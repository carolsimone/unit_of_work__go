package orchestration_uow

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	"unit_of_work__go/domain/k8s_deployment_status"
	"unit_of_work__go/domain/scheduler_status"
	"unit_of_work__go/infra/db"
)

func TestLoadAndRead(t *testing.T) {
	// Initialize the repo
	dbClient, err := db.NewDbConnection()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	// Create a new unit of work
	uow, errUoW := NewOrchestrationUnitOfWork(dbClient)
	if errUoW != nil {
		t.Fatalf("Failed to create unit of work: %v", errUoW)
	}
	ctx := context.Background()
	// Use the unit of work to run a transaction
	errTx := uow.RunInTx(ctx, func(ctx context.Context, tx db.Querier) error {
		// Load the scheduler status
		repoSS := uow.SchedulerStatusRepo()
		// Domain objects
		schedulerID := uuid.New()
		schedulerStatus := scheduler_status.SchedulerStatus{
			ID:        schedulerID,
			Name:      "daily",
			State:     scheduler_status.StateReady,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err := repoSS.Load(ctx, tx, &schedulerStatus)
		if err != nil {
			t.Fatalf("Failed to load SchedulerStatus domain object to DB: %v", err)
		}
		domainObject := k8s_deployment_status.K8sDeploymentStatus{
			ID:          uuid.New(),
			SchedulerID: schedulerID,
			State:       k8s_deployment_status.StatePending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		repoK8s := uow.K8sDeploymentRepo()
		err = repoK8s.Load(ctx, tx, &domainObject)
		return nil
	})
	if errTx != nil {
		t.Fatalf("Failed to run transaction in unit of work: %v", errTx)
	}
	// Read the scheduler status
	repoSS := uow.SchedulerStatusRepo()
	schedulerStatusRead, err := repoSS.Read(context.Background(), uow.DB())
	assert.Equal(t, 1, len(schedulerStatusRead))

	newCtx := context.Background()
	errTxDelete := uow.RunInTx(newCtx, func(newCtx context.Context, tx db.Querier) error {
		repoSS := uow.SchedulerStatusRepo()
		repoDS := uow.K8sDeploymentRepo()
		if errNewDS := repoDS.DeleteAll(newCtx, tx); errNewDS != nil {
			t.Fatalf("Failed to delete all k8s deployment statuses: %v", errNewDS)
		}
		if errNewSS := repoSS.DeleteAll(newCtx, tx); errNewSS != nil {
			t.Fatalf("Failed to delete all scheduler statuses: %v", errNewSS)
		}
		return nil
	})
	if errTxDelete != nil {
		t.Fatalf("Failed to run transaction for deletion in unit of work: %v", errTxDelete)
	}
	log.Println("Successfully deleted all statuses from the database")
}
