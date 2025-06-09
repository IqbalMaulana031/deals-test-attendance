package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"starter-go-gin/common/constant"
	"starter-go-gin/common/tools"
)

var (
	// ErrLoginNotFound is returned when the login is not found.
	ErrLoginNotFound = NewError(http.StatusNotFound, "Akun Login anda Belum terdaftar")
	// ErrLoginFailed login failed
	ErrLoginFailed = NewError(http.StatusUnauthorized, "Akun Login anda Salah")
	// ErrRecordNotFound represents error when record is not found.
	ErrRecordNotFound = NewError(http.StatusNotFound, "record not found")
	// ErrInternalServerError represents error when internal server error occurs.
	ErrInternalServerError = NewError(http.StatusInternalServerError, "internal server error")
	// ErrDuplicateReferenceID represents error when reference id is duplicate.
	ErrDuplicateReferenceID = NewError(http.StatusBadRequest, "reference id sudah ada")
	// ErrWrongLoginCredentials represents error when login credentials are wrong.
	ErrWrongLoginCredentials = NewError(http.StatusBadRequest, "email atau nomor HP salah")
	// ErrAccountNotRegistered represents error when account is not registered.
	ErrAccountNotRegistered = NewError(http.StatusBadRequest, "email atau nomor HP belum terdaftar")
	// ErrEmailNotRegistered represents error when email is not registered.
	ErrEmailNotRegistered = NewError(http.StatusBadRequest, "email tidak terdaftar")
	// ErrTooFarFromOffice represents error when user is too far from office.
	ErrTooFarFromOffice = NewError(http.StatusBadRequest, "anda terlalu jauh dari kantor")
	// ErrInvalidArgument represents error when invalid argument is passed.
	ErrInvalidArgument = NewError(http.StatusBadRequest, "invalid argument")
	// ErrWrongPasswordConfirmation define error if password confirmation is wrong
	ErrWrongPasswordConfirmation = NewError(http.StatusBadRequest, "Konfirmasi password kamu tidak sesuai.")
	// ErrOTPMismatch represents error when otp is mismatched.
	ErrOTPMismatch = NewError(http.StatusBadRequest, "Kode OTP Salah")
	// ErrUserNotFound represents error when user is not found.
	ErrUserNotFound = NewError(http.StatusBadRequest, "User tidak ditemukan")
	// ErrTokenNotValid represents error when token is not valid.
	ErrTokenNotValid = NewError(http.StatusBadRequest, "Token tidak valid")
	// ErrEmailAlreadyExist represents error when email is already exist.
	ErrEmailAlreadyExist = NewError(http.StatusBadRequest, "Email sudah terdaftar")
	// ErrPhoneNumberAlreadyExist represents error when phone number is already exist.
	ErrPhoneNumberAlreadyExist = NewError(http.StatusBadRequest, "Nomor HP sudah terdaftar")
	// ErrInvalidaTypeLogin represents error when type login is invalid.
	ErrInvalidaTypeLogin = NewError(http.StatusBadRequest, "akun anda tidak terdaftar disini")
	// ErrPaymentChannelNotFound represents error when payment channel is not found.
	ErrPaymentChannelNotFound = NewError(http.StatusBadRequest, "metode pembayaran tidak ditemukan")
	// ErrCustomerNotFound represents error when customer is not found.
	ErrCustomerNotFound = NewError(http.StatusBadRequest, "customer tidak ditemukan")
	// ErrProductQtyNotEnough represents error when product quantity is not enough.
	ErrProductQtyNotEnough = NewError(http.StatusBadRequest, "stok produk tidak mencukupi")
	// ErrOrderNotFound represents error when order is not found.
	ErrOrderNotFound = NewError(http.StatusBadRequest, "order tidak ditemukan")
	// ErrOrderCantBeCanceled represents error when order can't be canceled.
	ErrOrderCantBeCanceled = NewError(http.StatusBadRequest, "order tidak bisa dibatalkan")
	// ErrUnauthorizedDeleteEmployee represents error when employee can't be deleted.
	ErrUnauthorizedDeleteEmployee = NewError(http.StatusUnauthorized, "Anda tidak memiliki akses untuk menghapus pegawai")
	// ErrCustomerAlreadyExist represents error when customer is already exist.
	ErrCustomerAlreadyExist = NewError(http.StatusBadRequest, "Customer sudah terdaftar")
	// ErrProductAlreadyExist represents error when product is already exist.
	ErrProductAlreadyExist = NewError(http.StatusBadRequest, "Produk sudah ada")
	// ErrCanDeactivateCategory represents error when category can't be deactivated.
	ErrCanDeactivateCategory = NewError(http.StatusBadRequest, "Kategori tidak bisa dinonaktifkan")
	// ErrInvalidOTP represents error when otp is invalid.
	ErrInvalidOTP = NewError(http.StatusBadRequest, "Kode OTP tidak valid")
	// ErrOldEmailValidationOTP represents error when old email validation otp is invalid.
	ErrOldEmailValidationOTP = NewError(http.StatusBadRequest, "Validasi email lama tidak valid")
	// ErrExpiredOTP represents error when otp is expired.
	ErrExpiredOTP = NewError(http.StatusBadRequest, "Kode OTP kedaluwarsa")
	// ErrForbiddenDelete represents error when user can't delete data.
	ErrForbiddenDelete = NewError(http.StatusForbidden, "Anda tidak memiliki akses untuk menghapus data ini")
	// ErrPleaseChangePIN represents error when user must change pin.
	ErrPleaseChangePIN = NewError(http.StatusUnauthorized, "Silahkan ubah PIN, demi keamanan akun Anda")
	// ErrProductNotFound represents error when product is not found.
	ErrProductNotFound = NewError(http.StatusBadRequest, "Produk tidak ditemukan")
	// ErrCartNotFound represents error when cart is not found.
	ErrCartNotFound = NewError(http.StatusBadRequest, "Keranjang tidak ditemukan")
	// ErrUnauthorized represents error when user is unauthorized.
	ErrUnauthorized = NewError(http.StatusUnauthorized, "Unauthorized")
	// ErrForbidden represents error when user is forbidden.
	ErrForbidden = NewError(http.StatusForbidden, "Anda Tidak Memiliki Akses")
	// ErrUserAddressLimit represents error when user address limit is reached.
	ErrUserAddressLimit = NewError(http.StatusBadRequest, "Anda hanya bisa menambahkan 12 alamat")
	// ErrOrderCantBeConfirmed represents error when order can't be confirmed.
	ErrOrderCantBeConfirmed = NewError(http.StatusBadRequest, "order tidak bisa dikonfirmasi")
	// ErrOrderCantBeDelivered represents error when order can't be delivered.
	ErrOrderCantBeDelivered = NewError(http.StatusBadRequest, "order tidak bisa dikirim")
	// ErrOrderCantBeCompleted represents error when order can't be completed.
	ErrOrderCantBeCompleted = NewError(http.StatusBadRequest, "order tidak bisa diselesaikan")
	// ErrMerchantNotMatch represents error when merchant is not match.
	ErrMerchantNotMatch = NewError(http.StatusBadRequest, "Merchant tidak cocok")
	// ErrForbiddenAddress represents error when user can't access address.
	ErrForbiddenAddress = NewError(http.StatusForbidden, "Anda tidak memiliki akses untuk menggunakan alamat ini")
	// ErrPaymentChannelUnAvailable represents error when payment channel is unavailable.
	ErrPaymentChannelUnAvailable = NewError(http.StatusBadRequest, "metode pembayaran tidak tersedia")
	// ErrPaymentAlreadySuccess represents error when payment is already success.
	ErrPaymentAlreadySuccess = NewError(http.StatusBadRequest, "Pembayaran sudah berhasil")
	// MerchantCantBeReached represents error when merchant can't be reached.
	MerchantCantBeReached = NewError(http.StatusBadRequest, "Merchant Yang Anda Pilih Terlalu Jauh")
	// ErrShippingCostAlreadyExist represents error when shipping cost is already exist.
	ErrShippingCostAlreadyExist = NewError(http.StatusBadRequest, "Biaya pengiriman sudah ada")
	// ErrInvalidID represents error when id is invalid.
	ErrInvalidID = NewError(http.StatusBadRequest, "ID tidak valid")
	// ErrRecordAlreadyExists represents error when record is already exist.
	ErrRecordAlreadyExists = NewError(http.StatusBadRequest, "Data harga produk sudah ada")
	// ErrPromotionAlreadyExists represents error when promotion is already exist.
	ErrPromotionAlreadyExists = NewError(http.StatusBadRequest, "Promosi sudah ada")
	// ErrCantUpdateStatusApproval represents error when status approval can't be updated.
	ErrCantUpdateStatusApproval = NewError(http.StatusBadRequest, "Status persetujuan tidak bisa diubah")
	// ErrProvinceNotFound represents error when province is not found.
	ErrProvinceNotFound = NewError(http.StatusBadRequest, "Provinsi tidak ditemukan")
	// ErrRegencyNotFound represents error when regency is not found.
	ErrRegencyNotFound = NewError(http.StatusBadRequest, "Kota/Kabupaten tidak ditemukan")
	// ErrDistrictNotFound represents error when district is not found.
	ErrDistrictNotFound = NewError(http.StatusBadRequest, "Kecamatan tidak ditemukan")
	// ErrVillageNotFound represents error when village is not found.
	ErrVillageNotFound = NewError(http.StatusBadRequest, "Kelurahan tidak ditemukan")
	// ErrShippingTooFar represents error when shipping is too far.
	ErrShippingTooFar = NewError(http.StatusBadRequest, "Alamat pengiriman terlalu jauh")
	// ErrAccountNotApproved represents error when account is not approved.
	ErrAccountNotApproved = NewError(http.StatusBadRequest, "Akun Merchant Anda belum disetujui")
	// ErrVoucherLimitReached represents error when voucher limit is reached.
	ErrVoucherLimitReached = NewError(http.StatusBadRequest, "Voucher sudah mencapai batas maksimal")
	// ErrVoucherNotFound represents error when voucher is not found.
	ErrVoucherNotFound = NewError(http.StatusBadRequest, "Voucher tidak ditemukan")
	// ErrVoucherExpired represents error when voucher is expired.
	ErrVoucherExpired = NewError(http.StatusBadRequest, "Voucher sudah tidak berlaku")
	// ErrVoucherDayLimit represents error when voucher day limit is reached.
	ErrVoucherDayLimit = NewError(http.StatusBadRequest, "Voucher sudah mencapai limit harian")
	// ErrVoucherMonthLimit represents error when voucher month limit is reached.
	ErrVoucherMonthLimit = NewError(http.StatusBadRequest, "Voucher sudah mencapai limit bulanan")
	// ErrVoucherAlreadyUsed represents error when voucher is already used.
	ErrVoucherAlreadyUsed = NewError(http.StatusBadRequest, "Voucher sudah pernah digunakan")
	// ErrVoucherOutsideLocation represents error when voucher is outside location.
	ErrVoucherOutsideLocation = NewError(http.StatusBadRequest, "Voucher tidak berlaku di wilayah anda")
	// ErrCantUpdateDeliveryTime represents error when delivery time can't be updated.
	ErrCantUpdateDeliveryTime = NewError(http.StatusBadRequest, "Waktu pengiriman tidak bisa diubah")
	// ErrTypeOrder represents error when type order delivery is not match.
	ErrTypeOrder = NewError(http.StatusBadRequest, "Tipe order tidak sesuai")
	// ErrStatusOrder represents error when status order delivery is not match.
	ErrStatusOrder = NewError(http.StatusBadRequest, "Status order tidak sesuai")
	// ErrUpdateDeliveryTimeOnlyOnce represents error when delivery time can only be updated once.
	ErrUpdateDeliveryTimeOnlyOnce = NewError(http.StatusBadRequest, "Waktu pengiriman hanya bisa diubah 1 kali")
	// ErrNotFoundAddress represents error when address is not found.
	ErrNotFoundAddress = NewError(http.StatusBadRequest, "Alamat tidak ditemukan")
	// ErrCustomerNotRegisteredMyPertamina represents error when customer is not registered in MyPertamina.
	ErrCustomerNotRegisteredMyPertamina = NewError(http.StatusBadRequest, "Pelanggan belum terdaftar di MyPertamina")
	// ErrComplaintForDeliveredOrderOnly represents error when complaint is for delivered order only.
	ErrComplaintForDeliveredOrderOnly = NewError(http.StatusBadRequest, "Komplain hanya untuk order yang sudah dikirim")
	// ErrComplaintForComplainedOrderOnly represents error when complaint is for complained order only.
	ErrComplaintForComplainedOrderOnly = NewError(http.StatusBadRequest, "Komplain hanya untuk order yang sudah dikomplain")
	// ErrOrderAlreadyComplaint represents error when order is already complaint.
	ErrOrderAlreadyComplaint = NewError(http.StatusBadRequest, "Order sudah pernah dikomplain")
	// ErrComplaintNotFound represents error when complaint is not found.
	ErrComplaintNotFound = NewError(http.StatusBadRequest, "Komplain tidak ditemukan")
	// ErrOldPasswordMismatch represents error when old password is mismatch.
	ErrOldPasswordMismatch = NewError(http.StatusBadRequest, "PIN lama tidak sesuai")
	// ErrShippingImageAlreadyExists represents error when shipping image already exists
	ErrShippingImageAlreadyExists = NewError(http.StatusBadRequest, "Foto bukti sudah ada")
	// ErrShippingImageForDeliveredOrder represents error when shipping image is for delivered order only
	ErrShippingImageForDeliveredOrder = NewError(http.StatusBadRequest, "Foto bukti hanya untuk order yang dikirim atau sudah dikirim")
	// ErrAccountIsBroken represents error when account is broken
	ErrAccountIsBroken = NewError(http.StatusBadRequest, "Akun Anda sedang dalam perbaikan, silahkan hubungi Customer Service kami")
	// ErrBadRequest represents error when request is bad
	ErrBadRequest = NewError(http.StatusBadRequest, "Permintaan tidak valid")
	// ErrUnsuccessfulVerifyCaptcha represents error when captcha verification is unsuccessful
	ErrUnsuccessfulVerifyCaptcha = NewError(http.StatusBadRequest, "Verifikasi captcha gagal (waktu habis / duplikat)")
	// ErrScoreCaptchaLower represents error when score captcha is lower
	ErrScoreCaptchaLower = NewError(http.StatusBadRequest, "Nilai captcha terlalu rendah")
	// ErrMissMatchCaptcha represents error when captcha is mismatch
	ErrMissMatchCaptcha = NewError(http.StatusBadRequest, "Captcha tidak sesuai")
	// ErrUserAlreadyVoucherRedeemedToday represents error when user already redeemed voucher today
	ErrUserAlreadyVoucherRedeemedToday = NewError(http.StatusBadRequest, "Anda sudah menggunakan voucher hari ini")
	// ErrMerchantNotFound represents error when merchant is not found
	ErrMerchantNotFound = NewError(http.StatusBadRequest, "Merchant tidak ditemukan")
	// ErrMyPertaminaService represents error when MyPertamina service is not available
	ErrMyPertaminaService = NewError(http.StatusInternalServerError, "Gagal menghubungi MyPertamina")
	// ErrMIDNotSet represents error when MID is not set
	ErrMIDNotSet = NewError(http.StatusBadRequest, "MID belum diatur")
	// ErrVoucherNominalExceedProductPrice represents error when voucher nominal exceed product price
	ErrVoucherNominalExceedProductPrice = NewError(http.StatusBadRequest, "Nominal voucher melebihi harga produk")
	// ErrCannotDeleteUser represents error when user cannot be deleted
	ErrCannotDeleteUser = NewError(http.StatusBadRequest, "User tidak bisa dihapus")
	// ErrOrderInterval represents error when order interval is not match
	ErrOrderInterval = NewError(http.StatusBadRequest, "Interval order tidak sesuai")
	// ErrSubProductCategoryAlreadyExists represents error when sub product category already exists
	ErrSubProductCategoryAlreadyExists = NewError(http.StatusBadRequest, "Sub kategori produk sudah ada")
	// ErrProductCategoryNotFound represents error when product category is not found
	ErrProductCategoryNotFound = NewError(http.StatusBadRequest, "Kategori produk tidak ditemukan")
	// ErrProductOutOfStock represents error when product is out of stock
	ErrProductOutOfStock = NewError(http.StatusBadRequest, "Stok Produk sudah habis")
	// ErrComplaintFor135Order represents error when complaint is for 135 order only.
	ErrComplaintFor135Order = NewError(http.StatusBadRequest, "Pesanan 135 tidak bisa dikomplain")
	// ErrFileMaxSize represents error file max size
	ErrFileMaxSize = NewError(http.StatusBadRequest, "Ukuran file terlalu besar")
	// ErrCannotDeleteRole represents error when role cannot be deleted
	ErrCannotDeleteRole = NewError(http.StatusBadRequest, "Role tidak bisa dihapus, karena masih digunakan user lain")
)

// Error represents a data structure for error.
// It implements golang error interface.
type Error struct {
	// Code represents error code.
	Code int `json:"code"`
	// Message represents error message.
	// This is the message that exposed to the user.
	Message string `json:"message"`
}

// NewError creates an instance of Error.
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Error returns internal message in one string.
func (err *Error) Error() error {
	return fmt.Errorf("%d:%s", err.Code, err.Message)
}

// ParseError parses error message and returns an instance of Error.
func ParseError(err error) *Error {
	if err == nil {
		return nil
	}

	split := strings.Split(err.Error(), ":")

	fmt.Println(err)

	code, err := strconv.ParseInt(split[0], constant.Ten, constant.ThirtyTwo)

	if err != nil {
		return ErrInternalServerError
	}

	return NewError(int(code), split[1])
}

// ParseErrorValidation parses a validator.ValidationErrors into a more readable string format
func ParseErrorValidation(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return out
}

// ParseFieldError parses a validator.FieldError into a more readable string format
func parseFieldError(e validator.FieldError) string {
	// workaround to the fact that the `gt|gtfield=Start` gets passed as an entire tag for some reason
	// https://github.com/go-playground/validator/issues/926
	fieldPrefix := fmt.Sprintf("The field %s", tools.ToSnakeCase(e.Field()))
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "len":
		return fmt.Sprintf("%s must be %s characters long", fieldPrefix, e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldPrefix)
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", fieldPrefix, e.Param())
	case "isdefault":
		return fmt.Sprintf("%s must be %s", fieldPrefix, e.Param())
	default:
		// if it's a tag for which we don't have a good format string yet we'll try using the default english translator
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		}
		return fmt.Errorf("%v", e).Error()
	}
}

// ParseMarshallingError parses an unmarshalling error into a more readable string format
func parseMarshallingError(e json.UnmarshalTypeError) string {
	field := tools.ToSnakeCase(e.Field)
	return fmt.Sprintf("The field %s must be a %s", field, e.Type.String())
}
