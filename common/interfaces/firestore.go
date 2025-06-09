package interfaces

import (
	"context"
)

// FirestoreUseCase define interface for firestore
type FirestoreUseCase interface {
	Set(ctx context.Context, collection string, doc string, data *map[string]interface{}) error
	GetLocationInRadius(ctx context.Context, collection string, pointLat float64, pointLng float64, radiusInM int64) ([]map[string]interface{}, error)
	Get(ctx context.Context, collection string, doc string) (map[string]interface{}, error)
}
