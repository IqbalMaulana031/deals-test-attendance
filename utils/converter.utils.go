package utils

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"starter-go-gin/common/constant"
)

// StringToNullString convert string to sql null string
func StringToNullString(d string) sql.NullString {
	if d == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return sql.NullString{
		String: d,
		Valid:  true,
	}
}

// BoolToNullBool convert bool to sql null bool
func BoolToNullBool(d bool) sql.NullBool {
	return sql.NullBool{
		Bool:  d,
		Valid: true,
	}
}

// Float64ToNullFloat64 convert float64 to sql null float64
func Float64ToNullFloat64(d float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: d,
		Valid:   true,
	}
}

// Int32ToNullInt32 convert int32 to sql null int32
func Int32ToNullInt32(d int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: d,
		Valid: true,
	}
}

// Int64ToNullInt64 convert int64 to sql null int64
func Int64ToNullInt64(d int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: d,
		Valid: true,
	}
}

// TimeToNullTime convert time to sql null time
func TimeToNullTime(d time.Time) sql.NullTime {
	if d.IsZero() {
		return sql.NullTime{
			Time:  d,
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  d,
		Valid: true,
	}
}

// DateStringToTime convert date string to time
func DateStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}

	layout := "2006-01-02"
	t, errParse := time.Parse(layout, date)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t, nil
}

// DateTimeStringToTime convert date string to time
func DateTimeStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}

	layout := "2006-01-02 15:04:05"
	t, errParse := time.Parse(layout, date)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t, nil
}

// TimeStampsZStringToTime convert date string to time
func TimeStampsZStringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}

	layout := "2006-01-02T15:04:05Z07:00"
	t, errParse := time.Parse(layout, date)
	if errParse != nil {
		log.Println("errParse", errParse)
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t, nil
}

// ImageFullPath define image full path
func ImageFullPath(imgHost string, path string) string {
	return fmt.Sprintf("%s%s", imgHost, path)
}

// UUIDArrayToStringArray convert uuid array to string array
func UUIDArrayToStringArray(uuids []uuid.UUID) []string {
	result := make([]string, 0)
	for _, u := range uuids {
		result = append(result, fmt.Sprintf("'%s'", u))
	}

	return result
}

// RemoveDuplicateValues remove duplicate values from string array
func RemoveDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// DifferenceArray return difference between two string array
func DifferenceArray(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// NumberFormat format number
func NumberFormat(number int64) string {
	formatted := strconv.FormatInt(number, constant.Ten)
	return fmt.Sprintf("%s.%s", formatted[:len(formatted)-3], formatted[len(formatted)-3:])
}

// OperationDaysToBahasa convert comma separated operation days to bahasa
func OperationDaysToBahasa(operationDays string) string {
	days := strings.Split(operationDays, ",")
	var hari []string
	for _, day := range days {
		if day != "" {
			hari = append(hari, constant.ListDays[strings.ToLower(day)])
		}
	}
	return strings.Join(hari, ",")
}

// ConvertMerchantNameToCode convert merchant name to code
func ConvertMerchantNameToCode(merchantName string) string {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(merchantName, -1)

	result := ""
	for _, match := range matches {
		result += match
	}

	limitationResult := result
	if len(result) > constant.Twelve {
		limitationResult = result[0:12]
	}

	return limitationResult
}

// DateAndTimeStringToTime parses date and time string to time.Time
func DateAndTimeStringToTime(date string, timeStr string, tz *time.Location) (time.Time, error) {
	if date == "" || timeStr == "" {
		return time.Time{}, nil
	}

	layout := constant.DefaultTimeFormat
	t, errParse := time.ParseInLocation(layout, fmt.Sprintf("%s %s", date, timeStr), tz)
	if errParse != nil {
		return time.Time{}, fmt.Errorf("error while parsing date string to time : %v", errParse)
	}

	return t.In(tz), nil
}

func CurrentTimeInJakarta() time.Time {
	// Load lokasi Asia/Jakarta
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return time.Now() // fallback ke waktu default jika error
	}

	// Ambil waktu sekarang dalam zona waktu Asia/Jakarta
	currentTime := time.Now().In(location)
	return currentTime
}

func FirstAndLastDateOfCurrentMonth() (string, string) {
	// Ambil tanggal saat ini
	now := CurrentTimeInJakarta()

	// Tanggal pertama di bulan saat ini
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Tanggal terakhir di bulan saat ini
	lastDay := firstDay.AddDate(0, 1, -1) // Tambah 1 bulan, lalu kurangi 1 hari

	return firstDay.Format("2006-01-02"), lastDay.Format("2006-01-02")
}

func AddSevenHours(t time.Time) time.Time {
	return t.Add(7 * time.Hour)
}
