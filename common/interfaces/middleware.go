package interfaces

import (
	"context"
)

// MiddlewareTool is a service for middleware tool
type MiddlewareTool interface {
	RequiredPermission(ctx context.Context, permission string) (bool, error)
}
