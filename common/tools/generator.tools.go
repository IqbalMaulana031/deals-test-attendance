package tools

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	"starter-go-gin/common/constant"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// GenerateExternalID generate string with format prefix/YYYYMMDDHHmm/{random_int}
func GenerateExternalID(prefix string) string {
	timeNow := time.Now()
	res := prefix + fmt.Sprint(timeNow.Unix())

	return res
}

// ToSnakeCase convert string to snake case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// GenerateTrxID generate transaction ID
func GenerateTrxID(prefix string) string {
	res := prefix
	// rand.Seed(time.Now().UnixNano()) // #nosec
	// num := rand.Intn(constant.NinetyNineHundred-constant.Hundred) + constant.TenThousand // #nosec
	num := time.Now().UnixMilli()
	res = res + time.Now().Format("20060102") + fmt.Sprint(num)

	return res
}

// ReplaceString replace string
func ReplaceString(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// GenerateOTP generate OTP
func GenerateOTP(len int) string {
	var number string
	for i := 0; i < len; i++ {
		number += fmt.Sprint(rand.Intn(constant.Ten)) // #nosec
	}
	return number
}

// MappingStatus convert status from string to int
func MappingStatus(status string) string {
	switch status {
	case "pending":
		return "Pesanan Baru"
	case "confirmed":
		return "Pesanan Dikonfirmasi"
	case "delivered":
		return "Pesanan Dikirim"
	case "success":
		return "Pesanan Berhasil"
	case "failed":
		return "Pesanan Gagal"
	case "expired":
		return "Pesanan Kadaluwarsa"
	case "canceled":
		return "Pesanan Dibatalkan"
	default:
		return "Tidak Diketahui"
	}
}

// UUIDToString  convert uuid to string
func UUIDToString(id []uuid.UUID) []string {
	result := make([]string, 0)
	for _, v := range id {
		result = append(result, v.String())
	}
	return result
}

// StringToUUID convert string to uuid
func StringToUUID(id []string) []uuid.UUID {
	result := make([]uuid.UUID, 0)
	for _, v := range id {
		result = append(result, uuid.MustParse(v))
	}
	return result
}
