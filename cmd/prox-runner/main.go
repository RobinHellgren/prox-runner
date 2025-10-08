package main

import (
	githubHandler "github.com/RobinHellgren/prox-runner/v2/internal/webhookhandlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/github", githubHandler.HandleGitHubWebhook)
	r.Run()
}
