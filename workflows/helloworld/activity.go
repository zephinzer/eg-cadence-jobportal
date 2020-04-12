package main

import (
	"context"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

func Activity(ctx context.Context, value string) (string, error) {
	activity.GetLogger(ctx).Info("HelloWorld activity called", zap.String("value", value))
	return "processed: " + value, nil
}
