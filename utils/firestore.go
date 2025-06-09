package utils

import (
	"math"
	"strings"

	"starter-go-gin/common/constant"
)

var (
	// The meridional circumference of the earth in meters
	EarthMeriCircumference int64 = 40007860
	// Number of bits per geohash character
	BitsPerChar int = 5
	// Maximum length of a geohash in bits
	MaximumBitsPrecision int = 22 * BitsPerChar
	// Equatorial radius of the earth in meters
	EarthEqRadius float64 = 6378137.0
	// The following value assumes a polar radius of
	// const EARTH_POL_RADIUS = 6356752.3;
	// The formulate to calculate E2 is
	// E2 == (EarthEqRadius^2-EARTH_POL_RADIUS^2)/(EarthEqRadius^2)
	// The exact value is used here to avoid rounding errors
	E2 float64 = 0.00669447819799
	// Cutoff for rounding errors on double calculations
	Epsilon float64 = 1e-12
	// Length of a degree latitude at the equator
	MetersPerDegreeLatitude int64 = 110574
	// Characters used in location geohashes
	Base32 string = "0123456789bcdefghjkmnpqrstuvwxyz"
	// EarthRadiusInKm is the radius of the earth in kilometers
	EarthRadiusInKm float64 = 6371
)

// GeohashForLocation returns the geohash for the given location
func GeohashForLocation(latitude, longitude float64, precision int) string {
	if precision == 0 {
		precision = 10
	}
	latitudeRange := map[string]float64{"min": -constant.Ninety, "max": constant.Ninety}
	longitudeRange := map[string]float64{"min": -constant.OneHundredEighty, "max": constant.OneHundredEighty}
	hash := ""
	hashVal := 0
	bits := 0
	even := true
	for len(hash) < precision {
		val := latitude
		ranges := latitudeRange
		if even {
			val = longitude
			ranges = longitudeRange
		}

		mid := (ranges["min"] + ranges["max"]) / constant.Two
		if val > mid {
			hashVal = (hashVal << 1) + 1
			ranges["min"] = mid
		} else {
			hashVal = (hashVal << 1) + 0
			ranges["max"] = mid
		}
		even = !even
		if bits < constant.Four {
			bits++
		} else {
			bits = 0
			hash += string(Base32[hashVal])
			hashVal = 0
		}
	}
	return hash
}

// GeohashQueryBounds returns the bounding box query for the given location and radius
func GeohashQueryBounds(latitude, longitude, radius float64) [][]string {
	queryBits := int(math.Max(1, boundingBoxBits(latitude, radius)))
	geohashPrecision := int(math.Ceil(float64(queryBits) / float64(BitsPerChar)))
	coordinates := boundingBoxCoordinates(latitude, longitude, radius)
	queries := make([][]string, 0)
	for _, c := range coordinates {
		q := geohashQuery(GeohashForLocation(c[0], c[1], geohashPrecision), queryBits)
		queries = append(queries, q)
	}

	result := make([][]string, 0)
	// remove duplicates
	for index, query := range queries {
		duplicates := false
		// array.some
		for otherIndex, other := range queries {
			if index > otherIndex && query[0] == other[0] && query[1] == other[1] {
				duplicates = true
				break
			}
		}

		if !duplicates {
			result = append(result, query)
		}
	}

	return result
}

// DistanceBetween returns the distance between two locations, in kilometers.
func DistanceBetween(latitudeSrc, longitudeSrc, latitudeDst, longitudeDst float64) float64 {
	latDelta := degreesToRadians(latitudeDst - latitudeSrc)
	lonDelta := degreesToRadians(longitudeDst - longitudeSrc)
	a := (math.Sin(latDelta/constant.Two) * math.Sin(latDelta/constant.Two)) +
		(math.Cos(degreesToRadians(latitudeSrc)) * math.Cos(degreesToRadians(latitudeDst)) *
			math.Sin(lonDelta/constant.Two) * math.Sin(lonDelta/constant.Two))
	c := constant.Two * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadiusInKm * c
}

// =====================Private Function=====================

// geohashQuery returns a geohash query for the given geohash and precision
func geohashQuery(geohash string, bits int) []string {
	precision := int(math.Ceil(float64(bits) / float64(BitsPerChar)))
	if len(geohash) < precision {
		return []string{geohash, geohash + "~"}
	}
	geohash = geohash[:precision]
	base := geohash[:len(geohash)-1]

	lastValue := strings.Index(Base32, string(geohash[len(geohash)-1]))
	significantBits := bits - (len(base) * BitsPerChar)
	unusedBits := (BitsPerChar - significantBits)
	// delete unused bits
	startValue := (lastValue >> unusedBits) << unusedBits
	endValue := startValue + (1 << unusedBits)
	if endValue > constant.ThirtyOne {
		return []string{base + string(Base32[startValue]), base + "~"}
	} else {
		return []string{base + string(Base32[startValue]), base + string(Base32[endValue])}
	}
}

// boundingBoxCoordinates returns the coordinates of the bounding box for a given location and radius
func boundingBoxCoordinates(latitude, longitude, radius float64) [][]float64 {
	latDegrees := radius / float64(MetersPerDegreeLatitude)
	latitudeNorth := math.Min(constant.Ninety, latitude+latDegrees)
	latitudeSouth := math.Max(-constant.Ninety, latitude-latDegrees)
	longDegsNorth := metersToLongitudeDegrees(radius, latitudeNorth)
	longDegsSouth := metersToLongitudeDegrees(radius, latitudeSouth)
	longDegs := math.Max(longDegsNorth, longDegsSouth)

	return [][]float64{
		{latitude, longitude},
		{latitude, wrapLongitude(longitude - longDegs)},
		{latitude, wrapLongitude(longitude + longDegs)},
		{latitudeNorth, longitude},
		{latitudeNorth, wrapLongitude(longitude - longDegs)},
		{latitudeNorth, wrapLongitude(longitude + longDegs)},
		{latitudeSouth, longitude},
		{latitudeSouth, wrapLongitude(longitude - longDegs)},
		{latitudeSouth, wrapLongitude(longitude + longDegs)},
	}
}

// wrapLongitude wraps the longitude to [-180,180].
func wrapLongitude(longitude float64) float64 {
	if longitude <= constant.OneHundredEighty && longitude >= -constant.OneHundredEighty {
		return longitude
	}
	adjusted := longitude + constant.OneHundredEighty
	if adjusted > 0 {
		return math.Mod(adjusted, constant.ThreeHundredSixty) - constant.OneHundredEighty
	} else {
		return constant.OneHundredEighty - math.Mod(-adjusted, constant.ThreeHundredSixty)
	}
}

// boundingBoxBits returns the number of bits necessary to reach a given resolution, in meters, for the longitude at a given latitude.
func boundingBoxBits(latitude, size float64) float64 {
	latDeltaDegrees := size / float64(MetersPerDegreeLatitude)
	latitudeNorth := math.Min(constant.Ninety, latitude+latDeltaDegrees)
	latitudeSouth := math.Max(-constant.Ninety, latitude-latDeltaDegrees)
	bitsLat := math.Floor(latitudeBitsForResolution(size)) * constant.Two
	bitsLongNorth := math.Floor(longitudeBitsForResolution(size, latitudeNorth))*constant.Two - constant.One
	bitsLongSouth := math.Floor(longitudeBitsForResolution(size, latitudeSouth))*constant.Two - constant.One
	return findMinOfFloat(bitsLat, bitsLongNorth, bitsLongSouth, float64(MaximumBitsPrecision))
}

// latitudeBitsForResolution returns the number of bits necessary to reach a given resolution, in meters, for the latitude.
func latitudeBitsForResolution(resolution float64) float64 {
	return math.Min(math.Log2(float64(EarthMeriCircumference)/2/resolution), float64(MaximumBitsPrecision))
}

// longitudeBitsForResolution returns the number of bits necessary to reach a given resolution, in meters, for the longitude at a given latitude.
func longitudeBitsForResolution(resolution, latitude float64) float64 {
	degs := metersToLongitudeDegrees(resolution, latitude)
	result := float64(constant.One)
	if math.Abs(degs) > constant.ZeroPointZerosOne {
		result = math.Max(constant.One, math.Log2(constant.ThreeHundredSixty/degs))
	}
	return result
}

// metersToLongitudeDegrees returns the number of degrees longitude for a given distance, in meters, at a given latitude.
func metersToLongitudeDegrees(distance, latitude float64) float64 {
	radians := degreesToRadians(latitude)
	num := math.Cos(radians) * EarthEqRadius * math.Pi / constant.OneHundredEighty
	denom := constant.One / math.Sqrt(constant.One-E2*math.Sin(radians)*math.Sin(radians))
	deltaDeg := num * denom
	if deltaDeg < Epsilon {
		result := 0
		if distance > 0 {
			result = constant.ThreeHundredSixty
		}
		return float64(result)
	} else {
		return math.Min(constant.ThreeHundredSixty, distance/deltaDeg)
	}
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(degrees float64) float64 {
	return (degrees * math.Pi / constant.OneHundredEighty)
}

// findMinOfFloat returns the minimum value in a list of floats.
func findMinOfFloat(v ...float64) (m float64) {
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}
