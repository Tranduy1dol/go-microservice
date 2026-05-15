package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	ComponentHandler = "handler"
	ComponentUsecase = "usecase"
	ComponentAdapter = "adapter"
	ComponentAuth    = "auth"
	ComponentConfig  = "config"
	ComponentGRPC    = "adapter.grpc"
	ComponentMongo   = "adapter.mongo"
)

var root *slog.Logger

func Init(env string) {
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case slog.MessageKey:
					a.Key = "message"
				case slog.LevelKey:
					a.Key = "log.level"
				case slog.TimeKey:
					a.Key = "@timestamp"
				}

				return a
			},
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	root = slog.New(handler)
	slog.SetDefault(root)
}

func New(component string) *slog.Logger {
	if root == nil {
		Init("development")
	}

	return root.With("component", component)
}

func WithContext(ctx context.Context, component string) *slog.Logger {
	l := New(component)

	if reqID, ok := ctx.Value("request_id").(string); ok {
		l = l.With("request_id", reqID)
	}
	if userID, ok := ctx.Value("userID").(string); ok {
		l = l.With("user_id", userID)
	}

	return l
}
