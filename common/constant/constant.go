package constant

import (
	"strings"
)

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
	// SyncStatusPending define status pending sync
	SyncStatusPending = "pending"
	// SyncStatusProcess define status process sync
	SyncStatusProcess = "process"
	// SyncStatusCompleted define status completed sync
	SyncStatusCompleted = "completed"
	// SyncStatusFailed define status failed sync
	SyncStatusFailed = "failed"
	// Ascending define ascending constant
	Ascending = "asc"
	// Descending define descending constant
	Descending = "desc"
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
	// TenThousand define number ten thousand
	TenThousand = 10000
	// ThreeThousand define number three thousand
	ThreeThousand = 30000
	// Three define number three
	Three = 3
	// Seven define number seven
	Seven = 7
	// OneCommaFour define number one comma four
	OneCommaFour = 1.4
	// Ten define number ten
	Ten = 10
	// Five define number five
	Five = 5
	// Ninety define number ninety
	Ninety = 90
	// ThirtyTwo define number thirty two
	ThirtyTwo = 32
	// SixtyFour define number sixty four
	SixtyFour = 64
	// ContactUsEmail define contact us email
	ContactUsEmail = "adi.ardiansyah@gits.id"
	// CreatedAt define created at
	CreatedAt = "created_at"
	// UpdatedAt define updated at
	UpdatedAt = "updated_at"
	// Name define name
	Name = "name"
	// ProductIDCustom define
	ProductIDCustom = "SMEXPO"
	// Active define active
	Active = "active"
	// Inactive define inactive
	Inactive = "inactive"
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
	// PaymentGatewayMidtrans define payment gateway midtrans
	PaymentGatewayMidtrans = "midtrans"
	// PaymentGatewayXendit define payment gateway xendit
	PaymentGatewayXendit = "xendit"
	// OrderPrefix define order prefix
	OrderPrefix = "XXX"
	// PaymentStatusPending define payment status pending
	PaymentStatusPending = "pending"
	// PaymentStatusSuccess define payment status success
	PaymentStatusSuccess = "success"
	// CreatedBySystem define created by system
	CreatedBySystem = "system"
	// RoleMerchantOwner define role merchant owner
	RoleMerchantOwner = "owner"
	// RoleMerchantEmployeeAdmin define role merchant employee admin
	RoleMerchantEmployeeAdmin = "employee_admin"
	// RoleMerchantEmployeeCashier define role merchant employee cashier
	RoleMerchantEmployeeCashier = "employee_cashier"
	// RoleMerchantEmployeeDriver define role merchant employee driver
	RoleMerchantEmployeeDriver = "employee_driver"
	// LabelRoleMerchantEmployeeCashier define label role merchant employee cashier
	LabelRoleMerchantEmployeeCashier = "Kasir"
	// LabelRoleMerchantEmployeeAdmin define label role merchant employee admin
	LabelRoleMerchantEmployeeAdmin = "Admin"
	// CategoryTypeNonPertamina define category type non pertamina
	CategoryTypeNonPertamina = "non_pertamina"
	// CategoryNameNonPertamina define category name non pertamina
	CategoryNameNonPertamina = "Non Pertamina"
	// TwoHundred define number two hundred
	TwoHundred = 200
	// TwoHundredFifty define number two hundred fifty
	TwoHundredFifty = 250

	// DateDummy define date dummy
	DateDummy = "1970-01-01"
	// Keep define keep
	Keep = "keep"
	// Profit define profit
	Profit = "profit"
	// Loss define loss
	Loss = "loss"
	// Date format
	Date = "date"
	// SortByStatus define sort by status
	SortByStatus = "status"
	// SortByDate define sort by date
	SortByDate = "date"
	// ViewOrder define view order
	ViewOrder = "view_order"
	// RoleTypeCms define role type cms
	RoleTypeCms = "cms"
	// RedirectForgotPINLogin define redirect forgot pin login
	RedirectForgotPINLogin = "/login/forgot-pin"
	// FiveMinutesInSecond define five minutes in second
	FiveMinutesInSecond = 300
	// RedirectLogin define redirect login
	RedirectLogin = "/login"
	// Fifteen define number fifteen
	Fifteen = 15
	// One define number one
	One = 1
	// OneThousand define number one thousand
	OneThousand = 1000
	// RolePDSCustomer define role pds customer
	RolePDSCustomer = "customer"
	// EmailQA define email qa
	EmailQA = "qa-test@gits.id"
	// DefaultDate define default date
	DefaultDate = "1970-01-01"
	// OneWeekInHour define one week in hour
	OneWeekInHour = 168
	// RoleUser define role user
	RoleUser = "user"
	// RoleAdmin define role admin
	RoleAdmin = "admin"
	// DateFormat define date format
	DateFormat = "2006-01-02"
	// DateTransactionFormat define date transaction format
	DateTransactionFormat = "2006-01-02 15:04:05"
	// RoleTypePDS define role type pds
	RoleTypePDS = "pds"
	// RadiusNearbyMerchantInM define radius nearby merchant in meters
	RadiusNearbyMerchantInM = 5000
	// MerchantLocationFirestoreCollection define merchant location firestore collection
	// MerchantLocationFirestoreCollection = "merchant_locations"
	// DefaultMerchantOpenTime define default merchant open time
	DefaultMerchantOpenTime = "01:00:00"
	// DefaultMerchantCloseTime define default merchant close time
	DefaultMerchantCloseTime = "10:00:00"
	// DefaultMerchantCourierStartTime define default merchant courier start time
	DefaultMerchantCourierStartTime = "01:00:00"
	// DefaultMerchantCourierEndTime define default merchant courier end time
	DefaultMerchantCourierEndTime = "10:00:00"
	// TimeZoneAsiaJakarta define time zone asia jakarta
	TimeZoneAsiaJakarta = "Asia/Jakarta"
	// TimeZoneAsiaBangkok define time zone asia bangkok
	TimeZoneAsiaBangkok = "Asia/Bangkok"
	// SourcePlatformOrderMerchant define source platform order merchant
	SourcePlatformOrderMerchant = "merchant"
	// SourcePlatformOrderPDS define source platform order pds
	SourcePlatformOrderPDS = "pds"
	// SourcePlatformOrderWebPDS define source platform order web pds
	SourcePlatformOrderWebPDS = "web_pds"
	// SourcePlatformOrderMobilePDS define source platform order mobile pds
	SourcePlatformOrderMobilePDS = "mobile_pds"
	// SourcePlatformOrder135 define source platform order 135
	SourcePlatformOrder135 = "cms"
	// SourceMerchantOffline define source merchant offline
	SourceMerchantOffline = "offline"
	// TypeOrderCOD define type order cod
	TypeOrderCOD = "cod"
	// TypeOrderDelivery define type order delivery
	TypeOrderDelivery = "delivery"
	// TypeOrderPickup define type order pickup
	TypeOrderPickup = "pickup"
	// StatusOrderPDSPending define status order pds pending
	StatusOrderPDSPending = "pending"
	// StatusOrderPDSProcessing define status order pds processing
	StatusOrderPDSProcessing = "process"
	// StatusOrderPDSDelivery define status order pds delivery
	StatusOrderPDSDelivery = "delivered"
	// StatusOrderPDSCancelled define status order pds cancelled
	StatusOrderPDSCancelled = "canceled"
	// StatusOrderPDSCancel define status order pds cancel
	StatusOrderPDSCancel = "cancel"
	// StatusOrderPDSComplete define status order pds complete
	StatusOrderPDSComplete = "success"
	// StatusOrderPDSComplaint define status order pds complaint
	StatusOrderPDSComplaint = "complained"
	// NotificationTypeOrder define notification type order
	NotificationTypeOrder = "order"
	// ParamWithStock define param with stock
	ParamWithStock = "stock"
	// ParamWithProduct define param with product
	ParamWithProduct = "product"
	// PaymentNameCashOnDelivery define payment name cash on delivery
	PaymentNameCashOnDelivery = "Cash On Delivery"
	// Empty define empty
	Empty = "empty"
	// RoleSuperAdmin define role super admin
	RoleSuperAdmin = "super_admin"
	// MerchantApproved define merchant approved
	MerchantApproved = "approved"
	// CategoryPromoNational define category promo national
	CategoryPromoNational = "nasional"
	// CategoryPromoRegional define category promo regional
	CategoryPromoRegional = "regional"
	// FourHundred define number four hundred
	FourHundred = 400
	// FiveHundred define number five hundred
	FiveHundred = 500
	// ComplaintStatusCreated define complaint status created
	ComplaintStatusCreated = "created"
	// ComplaintStatusResendProduct define complaint status resend product
	ComplaintStatusResendProduct = "resend_product"
	// ComplaintStatusReceived define complaint status received
	ComplaintStatusReceived = "received_product"
	// ComplaintStatusCompleted define complaint status completed
	ComplaintStatusCompleted = "completed"
	// ComplaintStatusCanceled define complaint status canceled
	ComplaintStatusCanceled = "canceled"
	// Microsecond define microsecond
	Microsecond = 1000000
	// PromoCodeTypeGeneral define promo code type general
	PromoCodeTypeGeneral = "general"
	// PromoCodeTypeUnique define promo code type unique
	PromoCodeTypeUnique = "unique"
	// LogoPertamina define logo pertamina
	LogoPertamina = "/pertamina-files/Pertamina.png"
	// SettingSlugDiscountPromoMerchant define setting slug discount promo merchant
	SettingSlugDiscountPromoMerchant = "discount-promo-merchant"
	// DiscountTypePercentage define discount type percentage
	DiscountTypePercentage = "percentage"
	// DiscountTypeNominal define discount type nominal
	DiscountTypeNominal = "nominal"
	// SettingWhiteListMerchant define setting white list merchant
	SettingWhiteListMerchant = "whitelist-merchant"
	// SettingExpiredQR define setting expired qr
	SettingExpiredQR = "expired-qr-mypertamina"

	// merchant types
	MerchantTypeBrightStore = "bright-store"
	MerchantTypeBrightCafe  = "bright-cafe"

	// product pertamina bright gas
	ProductPertaminaGasFiveKG   = "Bright Gas 5,5 Kg"
	ProductPertaminaGasTwelveKG = "Bright Gas 12 Kg"
)

type ContextKey string

const UserIDKey ContextKey = "userID"

const (
	// default customer for order when customer_id not filed
	OrderDefaultCustomerNumber = "082117325080"
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

// Add these constants at the top of the file
const (
	EarthRadiusKm = 6371.0 // Radius of Earth in kilometers
	MetersToKm    = 1000.0
	KmPerDegLat   = 111.0 // km per degree latitude
	DegreesToRad  = 180.0 // for converting degrees to radians
	Epsilon       = 1e-10 // Small value to check for near-zero conditions
)

// ProductNameToProductID convert product name to product id
func ProductNameToProductID(productName string) string {
	productID := "Non-BBM"

	//nolint
	if strings.Contains(strings.ToLower(productName), "premium") {
		productID = "Premium"
	} else if strings.Contains(strings.ToLower(productName), "pertalite") {
		productID = "Pertalite"
	} else if strings.Contains(strings.ToLower(productName), "pertamax turbo") {
		productID = "Pertamax Turbo"
	} else if strings.Contains(strings.ToLower(productName), "pertamax") {
		productID = "Pertamax"
	} else if strings.Contains(strings.ToLower(productName), "dexlite") {
		productID = "Dexlite"
	} else if strings.Contains(strings.ToLower(productName), "pertamina dex") {
		productID = "Pertamina Dex"
	} else if strings.Contains(strings.ToLower(productName), "bio solar") {
		productID = "Bio Solar"
	} else if strings.Contains(strings.ToLower(productName), "elpiji") {
		productID = strings.ReplaceAll(strings.ToLower(productName), "elpiji", "LPG")
		productID = strings.ReplaceAll(productID, " ", "_")
		productID = strings.ToUpper(productID)
	} else if strings.Contains(strings.ToLower(productName), "bright gas") {
		productID = strings.ReplaceAll(strings.ToLower(productName), " ", "_")
		productID = strings.ReplaceAll(productID, ",", ".")
		productID = strings.ToUpper(productID)
	}

	return productID
}
func IsProductBrightGasFiveKG(productName string) bool {
	return strings.Contains(strings.ToLower(productName), strings.ToLower(ProductPertaminaGasFiveKG))
}

func IsProductBrightGasTwelveKG(productName string) bool {
	return strings.Contains(strings.ToLower(productName), strings.ToLower(ProductPertaminaGasTwelveKG))
}
