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
	// // SyncStatusPending define status pending sync
	// SyncStatusPending = "pending"
	// // SyncStatusProcess define status process sync
	// SyncStatusProcess = "process"
	// // SyncStatusCompleted define status completed sync
	// SyncStatusCompleted = "completed"
	// // SyncStatusFailed define status failed sync
	// SyncStatusFailed = "failed"
	// // ContactUsEmail define contact us email
	// ContactUsEmail = "admin@deals.id"
	// // CreatedAt define created at
	// CreatedAt = "created_at"
	// // UpdatedAt define updated at
	// UpdatedAt = "updated_at"
	// // Name define name
	// Name = "name"
	// // Active define active
	// Active = "active"
	// // Inactive define inactive
	// Inactive = "inactive"
	// // PaymentGatewayMidtrans define payment gateway midtrans
	// PaymentGatewayMidtrans = "midtrans"
	// // PaymentGatewayXendit define payment gateway xendit
	// PaymentGatewayXendit = "xendit"
	// // OrderPrefix define order prefix
	// OrderPrefix = "XXX"
	// // PaymentStatusPending define payment status pending
	// PaymentStatusPending = "pending"
	// // PaymentStatusSuccess define payment status success
	// PaymentStatusSuccess = "success"
	// // CreatedBySystem define created by system
	// CreatedBySystem = "system"
	// // LabelRoleMerchantEmployeeAdmin define label role merchant employee admin
	// LabelRoleMerchantEmployeeAdmin = "Admin"

	// // DateDummy define date dummy
	// DateDummy = "1970-01-01"
	// // Keep define keep
	// Keep = "keep"
	// // Profit define profit
	// Profit = "profit"
	// // Loss define loss
	// Loss = "loss"
	// // Date format
	// Date = "date"
	// // SortByStatus define sort by status
	// SortByStatus = "status"
	// // SortByDate define sort by date
	// SortByDate = "date"
	// // RoleTypeCms define role type cms
	// RoleTypeCms = "cms"
	// // RedirectForgotPINLogin define redirect forgot pin login
	// RedirectForgotPINLogin = "/login/forgot-pin"
	// // FiveMinutesInSecond define five minutes in second
	// FiveMinutesInSecond = 300
	// // RedirectLogin define redirect login
	// RedirectLogin = "/login"
	// // Fifteen define number fifteen
	// Fifteen = 15
	// // RolePDSCustomer define role pds customer
	// RolePDSCustomer = "customer"
	// // EmailQA define email qa
	// EmailQA = "qa-test@gits.id"
	// // DefaultDate define default date
	// DefaultDate = "1970-01-01"
	// // OneWeekInHour define one week in hour
	// OneWeekInHour = 168
	// // DateFormat define date format
	// DateFormat = "2006-01-02"
	// // DateTransactionFormat define date transaction format
	// DateTransactionFormat = "2006-01-02 15:04:05"
	// // RoleTypePDS define role type pds
	// RoleTypePDS = "pds"
	// // RadiusNearbyMerchantInM define radius nearby merchant in meters
	// RadiusNearbyMerchantInM = 5000
	// // MerchantLocationFirestoreCollection define merchant location firestore collection
	// // MerchantLocationFirestoreCollection = "merchant_locations"
	// // DefaultMerchantOpenTime define default merchant open time
	// DefaultMerchantOpenTime = "01:00:00"
	// // DefaultMerchantCloseTime define default merchant close time
	// DefaultMerchantCloseTime = "10:00:00"
	// // DefaultMerchantCourierStartTime define default merchant courier start time
	// DefaultMerchantCourierStartTime = "01:00:00"
	// // DefaultMerchantCourierEndTime define default merchant courier end time
	// DefaultMerchantCourierEndTime = "10:00:00"
	// // TimeZoneAsiaJakarta define time zone asia jakarta
	// TimeZoneAsiaJakarta = "Asia/Jakarta"
	// // TimeZoneAsiaBangkok define time zone asia bangkok
	// TimeZoneAsiaBangkok = "Asia/Bangkok"
	// // SourcePlatformOrderMerchant define source platform order merchant
	// SourcePlatformOrderMerchant = "merchant"
	// // SourcePlatformOrderPDS define source platform order pds
	// SourcePlatformOrderPDS = "pds"
	// // SourcePlatformOrderWebPDS define source platform order web pds
	// SourcePlatformOrderWebPDS = "web_pds"
	// // SourcePlatformOrderMobilePDS define source platform order mobile pds
	// SourcePlatformOrderMobilePDS = "mobile_pds"
	// // SourcePlatformOrder135 define source platform order 135
	// SourcePlatformOrder135 = "cms"
	// // SourceMerchantOffline define source merchant offline
	// SourceMerchantOffline = "offline"
	// // TypeOrderCOD define type order cod
	// TypeOrderCOD = "cod"
	// // TypeOrderDelivery define type order delivery
	// TypeOrderDelivery = "delivery"
	// // TypeOrderPickup define type order pickup
	// TypeOrderPickup = "pickup"
	// // StatusOrderPDSPending define status order pds pending
	// StatusOrderPDSPending = "pending"
	// // StatusOrderPDSProcessing define status order pds processing
	// StatusOrderPDSProcessing = "process"
	// // StatusOrderPDSDelivery define status order pds delivery
	// StatusOrderPDSDelivery = "delivered"
	// // StatusOrderPDSCancelled define status order pds cancelled
	// StatusOrderPDSCancelled = "canceled"
	// // StatusOrderPDSCancel define status order pds cancel
	// StatusOrderPDSCancel = "cancel"
	// // StatusOrderPDSComplete define status order pds complete
	// StatusOrderPDSComplete = "success"
	// // StatusOrderPDSComplaint define status order pds complaint
	// StatusOrderPDSComplaint = "complained"
	// // NotificationTypeOrder define notification type order
	// NotificationTypeOrder = "order"
	// // ParamWithStock define param with stock
	// ParamWithStock = "stock"
	// // ParamWithProduct define param with product
	// ParamWithProduct = "product"
	// // PaymentNameCashOnDelivery define payment name cash on delivery
	// PaymentNameCashOnDelivery = "Cash On Delivery"
	// // Empty define empty
	// Empty = "empty"
	// // RoleSuperAdmin define role super admin
	// RoleSuperAdmin = "super_admin"
	// // MerchantApproved define merchant approved
	// MerchantApproved = "approved"
	// // CategoryPromoNational define category promo national
	// CategoryPromoNational = "nasional"
	// // CategoryPromoRegional define category promo regional
	// CategoryPromoRegional = "regional"
	// // ComplaintStatusCreated define complaint status created
	// ComplaintStatusCreated = "created"
	// // ComplaintStatusResendProduct define complaint status resend product
	// ComplaintStatusResendProduct = "resend_product"
	// // ComplaintStatusReceived define complaint status received
	// ComplaintStatusReceived = "received_product"
	// // ComplaintStatusCompleted define complaint status completed
	// ComplaintStatusCompleted = "completed"
	// // ComplaintStatusCanceled define complaint status canceled
	// ComplaintStatusCanceled = "canceled"
	// // Microsecond define microsecond
	// Microsecond = 1000000
	// // PromoCodeTypeGeneral define promo code type general
	// PromoCodeTypeGeneral = "general"
	// // PromoCodeTypeUnique define promo code type unique
	// PromoCodeTypeUnique = "unique"
	// // LogoPertamina define logo pertamina
	// LogoPertamina = "/pertamina-files/Pertamina.png"
	// // SettingSlugDiscountPromoMerchant define setting slug discount promo merchant
	// SettingSlugDiscountPromoMerchant = "discount-promo-merchant"
	// // DiscountTypePercentage define discount type percentage
	// DiscountTypePercentage = "percentage"
	// // DiscountTypeNominal define discount type nominal
	// DiscountTypeNominal = "nominal"
	// // SettingWhiteListMerchant define setting white list merchant
	// SettingWhiteListMerchant = "whitelist-merchant"
	// // SettingExpiredQR define setting expired qr
	// SettingExpiredQR = "expired-qr-mypertamina"

	// // merchant types
	// MerchantTypeBrightStore = "bright-store"
	// MerchantTypeBrightCafe  = "bright-cafe"

	// // product pertamina bright gas
	// ProductPertaminaGasFiveKG   = "Bright Gas 5,5 Kg"
	// ProductPertaminaGasTwelveKG = "Bright Gas 12 Kg"

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
