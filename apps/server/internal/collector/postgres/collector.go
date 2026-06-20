package postgres

import (
	"bufio"
	"context"
	"io"
	"main/internal/repository/clickhouse"
	"os/exec"
)

type Collector struct {
	clickHouseRepo *clickhouse.Repository
}

func NewCollector(
	clickhouseRepo *clickhouse.Repository,
) *Collector {
	return &Collector{
		clickHouseRepo : clickhouseRepo,
	}
}



func (c *Collector) Start(ctx context.Context) error {
  scanner, err := ReadDockerLogs()
  if err != nil {
	return  err
  }

  for scanner.Scan(){
	line := scanner.Text()
	
  }
}

func ReadDockerLogs() (*bufio.Scanner, error) {
	cmd := exec.Command(
		"docker",
		"logs",
		"-f",
		"deployments-postgres-1",
	)

	sstdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	reader := io.MultiReader(sstdout, stderr)

	scanner := bufio.NewScanner(reader)

	return scanner, nil

}
