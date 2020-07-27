package util

import "time"

const (
	// DefaultServicePort is the port on which the svc is going to Listen and serve
	DefaultServicePort = 8888

	// Github client
	MaxRetries    = 8
	Max404Retries = 2
	MaxSleepTime  = 2 * time.Minute
	InitialDelay  = 2 * time.Second

	GithubBase = "https://github.com"

	AuthTokenKey  = "AUTHORIZATION"
	RepositoryKey = "Repository"

	Space = " "
)
