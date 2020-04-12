package main

import (
	"flag"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
)

func init() {
	activity.Register(Activity)
	workflow.Register(Workflow)
}

func main() {
	var mode string
	flag.StringVar(&mode, "m", "trigger", "Mode is worker or trigger.")
	flag.Parse()
	switch mode {
	case "worker":
		startWorker()
		select {}
	case "trigger":
		runTrigger()
	}
}
