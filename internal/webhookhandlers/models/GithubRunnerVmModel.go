package webhookhandlers

type GithubRunnerDto struct {
	RepositoryID   int64
	RepositoryName string
	RunnerID       *int64
	RunnerName     *string
	Label          string
	RequestedCpus  float64
	RequestedMem   int64
	RequestedDisk  int64
}
