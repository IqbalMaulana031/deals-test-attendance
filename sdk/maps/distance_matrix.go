package gmaps

import (
	"context"

	"github.com/pkg/errors"
	"googlemaps.github.io/maps"

	"starter-go-gin/config"
)

type GoogleMapsAPI struct {
	cfg config.Config
}

func NewGoogleAPI(cfg config.Config) *GoogleMapsAPI {
	return &GoogleMapsAPI{cfg: cfg}
}

func (g *GoogleMapsAPI) DistanceMatrix(origin, destination string) (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(g.cfg.Google.MapsAPI))
	if err != nil {
		return nil, errors.Wrap(err, "[GoogleMapsAPI-Distance Matrix] error get api key")
	}
	r := &maps.DistanceMatrixRequest{
		Origins:       []string{origin},
		Destinations:  []string{destination},
		Avoid:         maps.AvoidHighways,
		Mode:          maps.TravelModeDriving,
		TrafficModel:  maps.TrafficModelBestGuess,
		DepartureTime: "now",
	}
	result, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return nil, errors.Wrap(err, "[GoogleMapsAPI-Distance Matrix] error get distance matrix")
	}
	return result, nil
}
