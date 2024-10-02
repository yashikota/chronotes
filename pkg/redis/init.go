package redis

import "log/slog"

func Initialize() {
	_, err := Client.FlushAll(Ctx).Result()
	if err != nil {
		slog.Error("Failed to initialize Redis")
	}
	slog.Info("Successfully initialized Redis")
}
