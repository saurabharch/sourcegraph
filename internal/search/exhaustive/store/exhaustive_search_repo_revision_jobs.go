package store

import (
	"context"
	"time"

	"github.com/keegancsmith/sqlf"

	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/database/dbutil"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/search/exhaustive/types"
	dbworkerstore "github.com/sourcegraph/sourcegraph/internal/workerutil/dbworker/store"
)

var revSearchJobWorkerOpts = dbworkerstore.Options[*types.ExhaustiveSearchRepoRevisionJob]{
	Name:              "exhaustive_search_repo_revision_worker_store",
	TableName:         "exhaustive_search_repo_revision_jobs",
	ColumnExpressions: revSearchJobColumns,

	Scan: dbworkerstore.BuildWorkerScan(scanRevSearchJob),

	OrderByExpression: sqlf.Sprintf("exhaustive_search_repo_revision_jobs.state = 'errored', exhaustive_search_repo_revision_jobs.updated_at DESC"),

	StalledMaxAge: 60 * time.Second,
	MaxNumResets:  0,

	RetryAfter:    5 * time.Second,
	MaxNumRetries: 0,
}

// NewRevSearchJobWorkerStore returns a dbworkerstore.Store that wraps the "exhaustive_search_repo_revision_jobs" table.
func NewRevSearchJobWorkerStore(observationCtx *observation.Context, handle basestore.TransactableHandle) dbworkerstore.Store[*types.ExhaustiveSearchRepoRevisionJob] {
	return dbworkerstore.New(observationCtx, handle, revSearchJobWorkerOpts)
}

var revSearchJobColumns = []*sqlf.Query{
	sqlf.Sprintf("id"),
	sqlf.Sprintf("state"),
	sqlf.Sprintf("search_repo_job_id"),
	sqlf.Sprintf("revision"),
	sqlf.Sprintf("failure_message"),
	sqlf.Sprintf("started_at"),
	sqlf.Sprintf("finished_at"),
	sqlf.Sprintf("process_after"),
	sqlf.Sprintf("num_resets"),
	sqlf.Sprintf("num_failures"),
	sqlf.Sprintf("execution_logs"),
	sqlf.Sprintf("worker_hostname"),
	sqlf.Sprintf("cancel"),
	sqlf.Sprintf("created_at"),
	sqlf.Sprintf("updated_at"),
}

// ExhaustiveSearchRepoRevisionJobStore is the interface for interacting with "exhaustive_search_repo_revision_jobs".
type ExhaustiveSearchRepoRevisionJobStore interface {
	// CreateExhaustiveSearchRepoRevisionJob creates a new types.ExhaustiveSearchRepoRevisionJob.
	CreateExhaustiveSearchRepoRevisionJob(ctx context.Context, job types.ExhaustiveSearchRepoRevisionJob) (int64, error)
}

var _ ExhaustiveSearchRepoJobStore = &Store{}

func (s *Store) CreateExhaustiveSearchRepoRevisionJob(ctx context.Context, job types.ExhaustiveSearchRepoRevisionJob) (int64, error) {
	var err error
	ctx, _, endObservation := s.operations.createExhaustiveSearchRepoJob.With(ctx, &err, observation.Args{})
	defer endObservation(1, observation.Args{})

	row := s.Store.QueryRow(
		ctx,
		sqlf.Sprintf(createExhaustiveSearchRepoRevisionJobQueryFmtr, job.Revision, job.SearchRepoJobID),
	)

	var id int64
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

const createExhaustiveSearchRepoRevisionJobQueryFmtr = `
INSERT INTO exhaustive_search_repo_revision_jobs (revision, search_repo_job_id)
VALUES (%s, %s)
RETURNING id
`

func scanRevSearchJob(sc dbutil.Scanner) (*types.ExhaustiveSearchRepoRevisionJob, error) {
	var job types.ExhaustiveSearchRepoRevisionJob
	// required field for the sync worker, but
	// the value is thrown out here
	var executionLogs *[]any

	return &job, sc.Scan(
		&job.ID,
		&job.State,
		&job.SearchRepoJobID,
		&job.Revision,
		&dbutil.NullString{S: job.FailureMessage},
		&dbutil.NullTime{Time: &job.StartedAt},
		&dbutil.NullTime{Time: &job.FinishedAt},
		&dbutil.NullTime{Time: &job.ProcessAfter},
		&job.NumResets,
		&job.NumFailures,
		&executionLogs,
		&job.WorkerHostname,
		&job.Cancel,
		&job.CreatedAt,
		&job.UpdatedAt,
	)
}
