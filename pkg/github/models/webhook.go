package github

type WorkflowJobPayload struct {
	Action      WorkflowActionType `json:"action"`
	WorkflowJob WorkflowJob        `json:"workflow_job"`
	Repository  Repository         `json:"repository"`
}

type WorkflowJob struct {
	RunID      int64    `json:"run_id"`
	Labels     []string `json:"labels"`
	RunnerID   *int64   `json:"runner_id"`
	RunnerName *string  `json:"runner_name"`
}

type Repository struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type WorkflowActionType string

const (
	WorkflowRunCompletedActionType  WorkflowActionType = "completed"
	WorkflowRunQueuedActionType     WorkflowActionType = "queued"
	WorkflowRunInProgressActionType WorkflowActionType = "in_progress"
	WorkflowRunWaitingActionType    WorkflowActionType = "waiting"
)
