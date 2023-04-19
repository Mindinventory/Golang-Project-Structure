package database

import (
	"context"
	"time"
)

const (
	DefaultContextSeconds = 60
)

func DefaultTimeoutContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextSeconds*time.Second)
	return ctx, cancel
}

func CustomTimeoutContext(seconds int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	return ctx, cancel
}
