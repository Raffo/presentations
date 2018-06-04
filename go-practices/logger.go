package main

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "instance_id", 123)

	logger.Log("msg", "starting")
	NewWorker(log.With(logger, "component", "worker")).Run()
	NewSlacker(log.With(logger, "component", "slacker")).Run()
}
