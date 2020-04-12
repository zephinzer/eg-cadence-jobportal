package main

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/zap"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorker() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger.Info("creating new inbound transport channel")
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
	if err != nil {
		logger.Fatal("Failed to create transport channel", zap.Error(err))
	}
	logger.Info("creating new outbound transport channel")
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			ServiceName: {Unary: ch.NewSingleOutbound(BindAddress)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		logger.Fatal("Failed to create outbound transport channel: %v", zap.Error(err))
	}
	logger.Info("creating new workflow service client")
	service := workflowserviceclient.New(dispatcher.ClientConfig(ServiceName))

	workerOptions := worker.Options{
		MetricsScope: tally.NewTestScope(TaskListName, map[string]string{}),
		Logger:       logger,
	}
	worker := worker.New(service, DomainName, TaskListName, workerOptions)
	logger.Info("starting worker service")
	if err := worker.Start(); err != nil {
		logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}
