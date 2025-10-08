package webhookhandlers

import (
	"fmt"
	"strconv"
	"strings"

	githubHandlerModels "github.com/RobinHellgren/prox-runner/v2/internal/webhookhandlers/models"
	github "github.com/RobinHellgren/prox-runner/v2/pkg/github/models"
	"github.com/gin-gonic/gin"
)

func HandleGitHubWebhook(c *gin.Context) {
	var payload github.WorkflowJobPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	switch payload.Action {

	case github.WorkflowRunCompletedActionType:
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)

	case github.WorkflowRunWaitingActionType:
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)
	default:
		// Handle unknown action
	}
}

func CreateVmModelFromPayload(payload github.WorkflowJobPayload) githubHandlerModels.GithubRunnerDto {
	runnerParams := strings.Split(payload.WorkflowJob.Labels[0], "-")
	requestedCpus, err := strconv.ParseFloat(runnerParams[1], 64)
	if err != nil {
		fmt.Printf("Error parsing RequestedCpus: %v\n", err)
		requestedCpus = 0
	}
	requestedMem, err := strconv.ParseInt(runnerParams[2], 10, 64)
	if err != nil {
		fmt.Printf("Error parsing RequestedMem: %v\n", err)
		requestedMem = 0
	}
	requestedDisk, err := strconv.ParseInt(runnerParams[3], 10, 64)
	if err != nil {
		fmt.Printf("Error parsing RequestedDisk: %v\n", err)
		requestedDisk = 0
	}
	return githubHandlerModels.GithubRunnerDto{
		RepositoryID:   payload.Repository.ID,
		RepositoryName: payload.Repository.Name,
		RunnerID:       payload.WorkflowJob.RunnerID,
		RunnerName:     payload.WorkflowJob.RunnerName,
		Label:          runnerParams[0],
		RequestedCpus:  requestedCpus,
		RequestedMem:   requestedMem,
		RequestedDisk:  requestedDisk,
	}
}
