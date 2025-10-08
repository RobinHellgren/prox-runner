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
	fmt.Println("Received GitHub webhook")
	var payload github.WorkflowJobPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Printf("ERROR binding JSON: %v\n", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Action: %s\n", payload.Action)
	fmt.Printf("Repository: %s\n", payload.Repository.Name)
	fmt.Printf("Labels: %v\n", payload.WorkflowJob.Labels)

	switch payload.Action {

	case github.WorkflowRunCompletedActionType:
		fmt.Println("Handling COMPLETED action")
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)

	case github.WorkflowRunQueuedActionType:
		fmt.Println("Handling QUEUED action")
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)

	case github.WorkflowRunInProgressActionType:
		fmt.Println("Handling IN_PROGRESS action")
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)

	case github.WorkflowRunWaitingActionType:
		fmt.Println("Handling WAITING action")
		githubRunnerDto := CreateVmModelFromPayload(payload)
		fmt.Println(githubRunnerDto)
	default:
		fmt.Printf("Unknown action: %s\n", payload.Action)
	}
	c.JSON(200, gin.H{"status": "received"})
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
