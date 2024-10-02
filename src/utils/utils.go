package utils

import (
	"crypto/rand"
	"log/slog"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateUlid() string {
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, rand.Reader)
	if err != nil {
		slog.Error("Failed to generate ULID", "error", err)
		os.Exit(1)
	}
	return id.String()
}
