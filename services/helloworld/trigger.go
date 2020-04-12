package main

import (
	"context"
	"time"

	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
)

func runTrigger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("creating new inbound transport channel...")
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		logger.Fatal("failed to create transport channel", zap.Error(err))
	}
	logger.Debug("created new inbound transport channel...")

	logger.Info("creating new outbound transport channel...")
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			ServiceName: {Unary: ch.NewSingleOutbound(BindAddress)},
		},
	})
	logger.Info("starting outbound transport channel...")
	if err := dispatcher.Start(); err != nil {
		logger.Fatal("failed to create outbound transport channel: %v", zap.Error(err))
	}
	logger.Debug("created and started new outbound transport channel")

	logger.Info("creating new workflow service client...")
	service := workflowserviceclient.New(dispatcher.ClientConfig(ServiceName))
	logger.Debug("created new workflow service client")

	logger.Info("creating new workflow client...")
	workflowClient := client.NewClient(
		service,
		DomainName,
		&client.Options{
			Identity:     ServiceName,
			MetricsScope: tally.NoopScope,
		},
	)
	workflowOptions := client.StartWorkflowOptions{
		TaskList:                        TaskListName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	logger.Info("starting workflow...")
	if we, err := workflowClient.StartWorkflow(
		context.Background(),
		workflowOptions,
		Workflow,
		"cadence",
	); err != nil {
		logger.Fatal("failed to create workflow", zap.Error(err))
	} else {
		logger.Info("started Workflow", zap.String("WorkflowID", we.ID), zap.String("RunID", we.RunID))
	}
}
