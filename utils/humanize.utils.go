package utils

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Seconds-based time units
const (
	Day      = 24 * time.Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

// Time formats a time into a relative string.
//
// Time(someT) -> "3 weeks ago"
func Time(then time.Time) string {
	return RelTime(then, time.Now(), "yang lalu", "dari sekarang")
}

// A RelTimeMagnitude struct contains a relative time point at which
// the relative format of time will switch to a new format string.  A
// slice of these in ascending order by their "D" field is passed to
// CustomRelTime to format durations.
//
// The Format field is a string that may contain a "%s" which will be
// replaced with the appropriate signed label (e.g. "ago" or "from
// now") and a "%d" that will be replaced by the quantity.
//
// The DivBy field is the amount of time the time difference must be
// divided by in order to display correctly.
//
// e.g. if D is 2*time.Minute and you want to display "%d minutes %s"
// DivBy should be time.Minute so whatever the duration is will be
// expressed in minutes.
type RelTimeMagnitude struct {
	D      time.Duration
	Format string
	DivBy  time.Duration
}

var defaultMagnitudes = []RelTimeMagnitude{
	{time.Second, "sekarang", time.Second},
	{2 * time.Second, "1 detik %s", 1},
	{time.Minute, "%d detik %s", time.Second},
	{2 * time.Minute, "1 menit %s", 1},
	{time.Hour, "%d menit %s", time.Minute},
	{2 * time.Hour, "1 jam %s", 1},
	{Day, "%d jam %s", time.Hour},
	{2 * Day, "1 hari %s", 1},
	{Week, "%d hari %s", Day},
	{2 * Week, "1 minggu %s", 1},
	{Month, "%d minggu %s", Week},
	{2 * Month, "1 bulan %s", 1},
	{Year, "%d bulan %s", Month},
	{18 * Month, "1 tahun %s", 1},
	{2 * Year, "2 tahun %s", 1},
	{LongTime, "%d tahun %s", Year},
	{math.MaxInt64, "a long while %s", 1},
}

// RelTime formats a time into a relative string.
//
// It takes two times and two labels.  In addition to the generic time
// delta string (e.g. 5 minutes), the labels are used applied so that
// the label corresponding to the smaller time is applied.
//
// RelTime(timeInPast, timeInFuture, "earlier", "later") -> "3 weeks earlier"
func RelTime(a, b time.Time, albl, blbl string) string {
	return CustomRelTime(a, b, albl, blbl, defaultMagnitudes)
}

// CustomRelTime formats a time into a relative string.
//
// It takes two times two labels and a table of relative time formats.
// In addition to the generic time delta string (e.g. 5 minutes), the
// labels are used applied so that the label corresponding to the
// smaller time is applied.
func CustomRelTime(a, b time.Time, albl, blbl string, magnitudes []RelTimeMagnitude) string {
	lbl := albl
	diff := b.Sub(a)

	if a.After(b) {
		lbl = blbl
		diff = a.Sub(b)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].D > diff
	})

	if n >= len(magnitudes) {
		n = len(magnitudes) - 1
	}
	mag := magnitudes[n]
	args := []interface{}{}
	escaped := false
	for _, ch := range mag.Format {
		if escaped {
			switch ch {
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.DivBy)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.Format, args...)
}

// FormatRupiah to convert number 10000 to Rp 10.000
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.ReplaceAll(humanizeValue, ",", ".")
	return "Rp " + stringValue
}

// FormatRupiahWithoutRp to convert number 10000 to 10.000
func FormatRupiahWithoutRp(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.ReplaceAll(humanizeValue, ",", ".")
	return stringValue
}
