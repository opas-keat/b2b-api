package logs

import (
	"logs/repo/logs"
)

type ServiceImpl struct {
	logsRepo logs.Repo
}

func New(
	logsRepo logs.Repo,
) (Service, error) {
	return &ServiceImpl{
		logsRepo,
	}, nil
}
