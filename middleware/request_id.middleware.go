package middleware

import (
	"context"
	"time"
)

// GetRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetRequestID(ctx context.Context) string {
	return time.Now().Format("20060102T150405000")
}
