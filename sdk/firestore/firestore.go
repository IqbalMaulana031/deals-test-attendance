package firestore

import (
	"context"

	"cloud.google.com/go/firestore"

	"starter-go-gin/common/constant"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/utils"
)

// Firestore is an struct for Firestore SDK
type Firestore struct {
	cfg    config.Config
	client *firestore.Client
}

// NewFirestore initiate Firestore SDK
func NewFirestore(cfg config.Config) *Firestore {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.Google.ProjectID)
	if err != nil {
		logger.Error(ctx, err)
	}

	return &Firestore{
		cfg:    cfg,
		client: client,
	}
}

// Set data to firestore
func (f *Firestore) Set(ctx context.Context, collection string, doc string, data *map[string]interface{}) error {
	_, err := f.client.Collection(collection).Doc(doc).Set(ctx, data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// Get data from firestore
func (f *Firestore) Get(ctx context.Context, collection string, doc string) (map[string]interface{}, error) {
	snapshot, err := f.client.Collection(collection).Doc(doc).Get(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	var data map[string]interface{}
	err = snapshot.DataTo(&data)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return data, nil
}

// GetLocationInRadius get location in radius (using geohash)
func (f *Firestore) GetLocationInRadius(ctx context.Context, collection string, pointLat float64, pointLng float64, radiusInM int64) ([]map[string]interface{}, error) {
	bounds := utils.GeohashQueryBounds(pointLat, pointLng, float64(radiusInM))
	snapshots := make([]*firestore.DocumentSnapshot, 0)
	for _, b := range bounds {
		s, err := f.client.Collection(collection).
			OrderBy("geohash", firestore.Asc).
			StartAt(b[0]).
			EndAt(b[1]).
			Documents(ctx).
			GetAll()

		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		snapshots = append(snapshots, s...)
	}

	result := make([]map[string]interface{}, 0)
	for _, s := range snapshots {
		var data map[string]interface{}
		err := s.DataTo(&data)
		if err != nil {
			logger.Error(ctx, err)
			continue
		}
		lat := data["latitude"].(float64)
		lng := data["longitude"].(float64)

		// mapsAPI := google_maps_api.NewGoogleAPI(f.cfg)
		// origin := fmt.Sprintf("%v,%v", lat, lng)
		// destination := fmt.Sprintf("%v,%v", pointLat, pointLng)
		// resultAPI, err := mapsAPI.DistanceMatrix(origin, destination)
		// if err != nil {
		// 	logger.ErrorWithStr(ctx, "[FirestoreSDK-GetLocationInRadius] Error when get distance", err)
		// 	continue
		// }

		// distance := resultAPI.Rows[0].Elements[0].Distance.Meters
		// distanceInKm := float64(distance) / 1000

		distanceInKm := utils.DistanceBetween(lat, lng, pointLat, pointLng)
		distanceInM := distanceInKm * constant.Thousand
		if distanceInM <= float64(radiusInM) {
			r := map[string]interface{}{
				"id":          s.Ref.ID,
				"latitude":    lat,
				"longitude":   lng,
				"geohash":     data["geohash"],
				"distance_km": distanceInKm,
			}
			result = append(result, r)
		}
	}

	return result, nil
}
