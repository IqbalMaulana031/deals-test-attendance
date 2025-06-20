package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"

	"starter-go-gin/common/constant"
)

const (
	four = 4
)

// IPLocation define API Response for GetGeo API
type IPLocation struct {
	IP       string    `json:"ip"`
	Type     string    `json:"type"`
	Location *Location `json:"location"`
	Area     *Area     `json:"area"`
	Asn      *Asn      `json:"asn"`
	City     *City     `json:"city"`
	Status   string    `json:"status"`
}

// Location define API Response for GetGeo API
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Area define API Response for GetGeo API
type Area struct {
	Code      string `json:"code"`
	Geonameid int    `json:"geonameid"`
	Name      string `json:"name"`
}

// Asn define API Response for GetGeo API
type Asn struct {
	Number       int    `json:"number"`
	Organisation string `json:"organisation"`
}

// City define API Response for GetGeo API
type City struct {
	Geonameid  int    `json:"geonameid"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

// GetLocationByIP get geolocation by user IP
func GetLocationByIP(ip string) (*IPLocation, error) {
	var headers []CallerHeader

	data := &IPLocation{}

	splitIP := strings.Split(ip, ".")

	if len(splitIP) > four {
		return nil, nil
	}

	if ip == "" {
		return nil, nil
	}

	payload := map[string]string{
		"api_key":  os.Getenv("IP_GEO_KEY"),
		"format":   "json",
		"filter":   "city,area,asn",
		"language": "en",
	}

	res, err := CallAPI(http.MethodGet, fmt.Sprintf("%s/%s", os.Getenv("IP_GEO_URL"), ip), headers, nil, payload)
	if err != nil {
		return nil, errors.Wrap(err, "[Utils-GetLocationByAPI-CallAPI]")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			// Log or handle the error appropriately
			fmt.Println("Error closing response body:", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, nil
	}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// CalculateTotalPoints menghitung total poin berdasarkan jumlah Rupiah.
// Setiap kelipatan 10.000 Rupiah memberikan 5 poin, dengan maksimal 250 poin.
func CalculateTotalPoints(amount int64) int64 {
	// Hitung kelipatan 10.000 Rupiah
	multiples := amount / constant.TenThousand

	// Hitung poin
	points := multiples * constant.Five

	// Pastikan poin tidak melebihi 250
	if points > constant.TwoHundredFifty {
		return constant.TwoHundredFifty
	}

	return points
}
