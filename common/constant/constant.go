package constant

var (
	ListStatusOrder      = []string{"pending", "confirmed", "delivered", "complained", "success", "failed", "canceled"}
	ListLabelStatusOrder = []string{"Pesanan Baru", "Dikonfirmasi", "Dikirim", "Dikomplain", "Selesai", "Gagal", "Batal"}

	ListStatusOrderPDS      = []string{"process", "delivered", "success", "canceled", "complained"}
	ListLabelStatusOrderPDS = []string{"Diproses", "Dikirim", "Selesai", "Dibatalkan", "Dikomplain"}

	ListPopularEmail = []string{"gmail", "yahoo", "hotmail", "rocketmail", "msn", "icloud", "live", "outlook", "facebook"}

	MapListStatusLinkaja = map[string]string{
		"PGPINQ":              "unpaid",
		"FAILED":              "failed",
		"LANDINGPAGE":         "unpaid",
		"AUTHENTICATION":      "unpaid",
		"SUCCESS_COMPLETED":   "success",
		"CANCELLEDBYCUSTOMER": "failed",
		"SUCCESS_REVERSED":    "success",
	}

	MapListPaymentStatus = map[string]string{
		"unpaid":  "Menunggu Pembayaran",
		"success": "Pembayaran Berhasil",
		"failed":  "Pembayaran Gagal",
	}

	ListDays = map[string]string{
		"monday":    "Senin",
		"tuesday":   "Selasa",
		"wednesday": "Rabu",
		"thursday":  "Kamis",
		"friday":    "Jumat",
		"saturday":  "Sabtu",
		"sunday":    "Minggu",
	}
)

const (
	LinkAjaEDDTerminalId = "testing_linkaja_edd" //nolint:stylecheck
	LinkAjaEDDUserKey    = "wcotest1091"
	LinkAjaEDDPassword   = "@wcotest12"
	LinkAjaEDDSignature  = "wcotestsign"

	LinkAjaTerminalId = "testing_linkaja_wco" //nolint:stylecheck
	LinkAjaUserKey    = "wcotest1091"
	LinkAjaPassword   = "@wcotest12"
	LinkAjaSignature  = "wcotestsign"
)

const (
	// DefaultLimit default limit for pagination
	DefaultLimit = 10
	// ZeroPointZerosOne define number zero point zeros one
	ZeroPointZerosOne = 0.000001
	// Four define number four
	Four = 4
	// Two define number two
	Two = 2
	// Sixty define number sixty
	Sixty = 60
	// Six define number six
	Six = 6
	// OneHundredEighty define number one hundred eighty
	OneHundredEighty = 180
	// OneHundreds
	OneHundreds = 100
	// TwentyFourHour define twenty-four hours
	TwentyFourHour = 24
	// TwentySix define number twenty six
	TwentySix = 26
	// Twelve define number twelve
	Twelve = 12
	// DaysInOneYear define days in one year
	DaysInOneYear = 365
	// Thirty define number thirty
	Thirty = 30
	// FortyFive
	FortyFive = 45
	// ThirtyOne define number thirty one
	ThirtyOne = 31
	// Fifty define number fifty
	Fifty = 50
	// Hundred define number hundred
	Hundred = 100
	// ThreeHundred define number ThreeHundred
	ThreeHundred = 300
	// ThreeHundredSixty define number ThreeHundredSixty
	ThreeHundredSixty = 360
	// Thousand define number thousand
	Thousand = 1000
	// NinetyNineHundred define number ninety nine hundred
	NinetyNineHundred = 9999
	// ThreeThousand define number three thousand
	ThreeThousand = 30000
	// Three define number three
	Three = 3
	// Seven define number seven
	Seven = 7
	// OneCommaFour define number one comma four
	OneCommaFour = 1.4
	// Five define number five
	Five = 5
	// Ninety define number ninety
	Ninety = 90
	// SixtyFour define number sixty four
	SixtyFour = 64
	// Ten define number ten
	Ten = 10
	// ThirtyTwo define number thirty two
	ThirtyTwo = 32
	// FourHundred define number four hundred
	FourHundred = 400
	// FiveHundred define number five hundred
	FiveHundred = 500
	// TenThousand define number ten thousand
	TenThousand = 10000
	// One define number one
	One = 1
	// OneThousand define number one thousand
	OneThousand = 1000
	// TwoHundred define number two hundred
	TwoHundred = 200
	// TwoHundredFifty define number two hundred fifty
	TwoHundredFifty = 250
	// employeeRoleName define employee role name
	EmployeeRoleName = "employee"
	// Ascending define ascending constant
	Ascending = "asc"
	// Descending define descending constant
	Descending = "desc"
	// DefaultTimeFormat default time format
	DefaultTimeFormat = "2006-01-02 15:04:05"
	// DefaultTimeFormatShort default time format short
	DefaultTimeFormatShort = "2006-01-02"
	// DefaultTimeFormatJustTime default time format just time
	DefaultTimeFormatJustTime = "15:04:05"
	// DefaultTimeFormatJustTimeShort default time format just time short
	DefaultTimeFormatJustTimeShort = "15:04"
	// MethodPOST define method post
	MethodPOST = "POST"
	// MethodGET define method get
	MethodGET = "GET"
	// RoleAdmin define role admin
	RoleAdmin = "admin"
)

const (
	// RequestHeaderName request header name
	RequestHeaderName string = "X-Name"
	// RequestHeaderUserID request header user id
	RequestHeaderUserID string = "X-User-ID"
	// RequestHeaderRoleID request header role id
	RequestHeaderRoleID string = "X-Role-ID"
	// RequestHeaderPermission request header permission
	RequestHeaderPermission string = "X-Permission"
	// RequestHeaderMerchantID request header merchant id
	RequestHeaderMerchantID string = "X-Merchant-ID"
	// RequestHeaderJTI request header jti
	RequestHeaderJTI string = "X-JTI"
)

type ContextKey string

const UserIDKey ContextKey = "userID"
