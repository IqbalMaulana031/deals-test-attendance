package utils

import (
	"math"

	"github.com/alicebob/miniredis/v2/geohash"

	"starter-go-gin/common/constant"
)

// CalculateDistance calculates distance between two points
func CalculateDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793
	const statuteMiles float64 = 1.1515

	radlat1 := PI * lat1 / constant.OneHundredEighty
	radlat2 := PI * lat2 / constant.OneHundredEighty

	theta := lng1 - lng2
	radtheta := PI * theta / constant.OneHundredEighty

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * constant.OneHundredEighty / PI
	dist = dist * constant.Sixty * statuteMiles

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist *= 1.609344
		} else if unit[0] == "N" {
			dist *= 0.8684
		}
	}

	return dist
}

func Geohash(lat float64, lng float64) string {
	geohash := geohash.Encode(lat, lng)
	return geohash
}
