package domain

import (
	"context"
	"log/slog"
	"time"
)

func GetData(ctx context.Context) string {
	if logger, ok := ctx.Value("logger").(*slog.Logger); ok {
		logger.InfoContext(ctx, "GetData running")
	} else {
		slog.InfoContext(ctx, "GetData running")
	}
	time.Sleep(10 * time.Millisecond)
	return "data"
}
