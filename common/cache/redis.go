package cache

const (
	// FiveMinutes is ttl for the cache
	FiveMinutes = 300
	// OneHour is ttl for the cache
	OneHour = 3600
	// OneMonth is ttl for the cache
	OneMonth = 2592000
	// OneWeek is ttl for the cache
	OneWeek = 604800
	// OneYear is ttl for the cache
	OneYear = 31536000
	// FourMonths is ttl for the cache
	FourMonths = 12960000
	// OneDay is ttl for the cache
	OneDay = 86400
)

const (
	prefix = "deals-api"
	// UserRoleByUserID is a redis key for find cms user role by user id.
	UserRoleByUserID = prefix + ":user-role:find-by-user-id:%v"
	// // PermissionFindByName is a redis key for find cms permission by name.
	// PermissionFindByName = prefix + ":permission:find-by-name:%v"
	// // RolePermissionFindByRoleIDAndPermissionID is a redis key for find cms role permission by role id and permission id.
	// RolePermissionFindByRoleIDAndPermissionID = prefix + ":role-permission:find-by-role-id-and-permission-id:%v:%v"
	// RoleFindByID is a redis key for find cms role by id.
	RoleFindByID = prefix + ":role:find-by-id:%v"
	// // BankFindAllTotal is a redis key for find all bank total.
	// BankFindAllTotal = prefix + ":bank:find-all-total:%v:%v:%v:%v:%v"
	// // BankFindAll is a redis key for find all bank.
	// BankFindAll = prefix + ":bank:find-all:%v:%v:%v:%v:%v"
	// // RolePermissionFindByRoleID is a redis key for find cms role permission by role id.
	// RolePermissionFindByRoleID = prefix + ":role-permission:find-by-role-id:%v"
	// UserByID is a redis key for find cms user by id.
	UserByID = prefix + ":user:find-user-by-id:%v"
	// ShiftByID is a redis key for find cms shift by id.
	ShiftByID = prefix + ":shift:find-by-id:%v"
	// ShiftAndDetailsByID is a redis key for find cms shift and details by id.
	ShiftAndDetailsByID = prefix + ":shift:find-and-details-by-id:%v"
	// shiftDetailByID is a redis key for find cms shift detail by id.
	ShiftDetailByID = prefix + ":shift-detail:find-by-id:%v"
	// AttendanceByID is a redis key for find attendance by id.
	AttendanceByID = prefix + ":attendance:find-by-id:%v"

	// // ProductCategoryFindAllTotal is a redis key for find all product category total.
	// ProductCategoryFindAllTotal = prefix + ":product-category:find-all-total"
	// // ProductCategoryFindAll is a redis key for find all product category.
	// ProductCategoryFindAll = prefix + ":product-category:find-all"
	// // RoleEmployeeFindAll is a redis key for find all role employee.
	// RoleEmployeeFindAll = prefix + ":role-employee:find-all"
	// // RolePermissionFindByRoleIDWithAssociation is a redis key for find cms role permission by role id with association.
	// RolePermissionFindByRoleIDWithAssociation = prefix + ":role-permission:find-by-role-id-with-association:%v"
	// // RolePermissionFindMenuIDByRoleIDs is a redis key for find cms role permission menu id by role ids.
	// RolePermissionFindMenuIDByRoleIDs = prefix + ":role-permission:find-menu-id-by-role-ids:%v"
	// // RolePermissionFindByRoleIDAndMenuID is a redis key for find cms role permission by role id and menu id.
	// RolePermissionFindByRoleIDAndMenuID = prefix + ":role-permission:find-by-role-id-and-menu-id:%v:%v"
	// // PertaminaProductFindAllTotal is a redis key for find all pertamina product total.
	// PertaminaProductFindAllTotal = prefix + ":pertamina-product:find-all-total:%v:"
	// // PertaminaProductFindAll is a redis key for find all pertamina product.
	// PertaminaProductFindAll = prefix + ":pertamina-product:find-all:%v:"
	// // MerchantProductCategoryTotalFindByMerchantID is a redis key for find total merchant product category by merchant id.
	// MerchantProductCategoryTotalFindByMerchantID = prefix + ":merchant-product-category-total:find-by-merchant-id:%v"
	// // MerchantProductCategoryFindByMerchantID is a redis key for find all product category.
	// MerchantProductCategoryFindByMerchantID = prefix + ":merchant-product-category:find-by-merchant-id:%v"
	// // MerchantProductCategoryTotalFindActiveByMerchantID is a redis key for find total active merchant product category by merchant id.
	// MerchantProductCategoryTotalFindActiveByMerchantID = prefix + ":merchant-product-category-total:find-active-by-merchant-id:%v"
	// // MerchantProductCategoryFindActiveByMerchantID is a redis key for find all active product category.
	// MerchantProductCategoryFindActiveByMerchantID = prefix + ":merchant-product-category:find-active-by-merchant-id:%v"
	// // MerchantProductCategoryFindByMerchantIDAndProductCategoryID is a redis key for find product category by merchant id and product category id.
	// MerchantProductCategoryFindByMerchantIDAndProductCategoryID = prefix + ":merchant-product-category:find-by-merchant-id-and-product-category-id:%v:%v"
	// // CustomerFindAllTotal is a redis key for find all customer total.
	// CustomerFindAllTotal = prefix + ":customer:find-all-total:%v:%v:%v:%v:%v:%v:%v"
	// // CustomerFindAll is a redis key for find all customer.
	// CustomerFindAll = prefix + ":customer:find-all:%v:%v:%v:%v:%v:%v:%v"
	// // CustomerFindByPhoneNumber is a redis key for find customer by phone number.
	// CustomerFindByPhoneNumber = prefix + ":customer:find-by-phone-number:%v:%v"
	// // CustomerFindByPhoneNumber is a redis key for find customer by phone number without merchant id.
	// CustomerFindByPhoneNumberV2 = prefix + ":customer:find-by-phone-number-standalone:%v"
	// // CustomerFindByID is a redis key for find customer by id.
	// CustomerFindByID = prefix + ":customer:find-by-id:%v"
	// // BusinessFieldFindAll is a redis key for find all business field.
	// BusinessFieldFindAll = prefix + ":business-field:find-all"
	// // BusinessTypeFindAll is a redis key for find all business type.
	// BusinessTypeFindAll = prefix + ":business-type:find-all"
	// // DistrictFindByRegencyID is a redis key for find regency by id.
	// DistrictFindByRegencyID = prefix + ":regency:find-by-id:%v:%v"
	// // ProvinceFindByID is a redis key for find province by id.
	// ProvinceFindByID = prefix + ":province:find-by-id:%v"
	// // ProvinceFindAll is a redis key for find all province.
	// ProvinceFindAll = prefix + ":province:find-all:%v"
	// // ProvinceFindByName is a redis key for find province by name.
	// ProvinceFindByName = prefix + ":province:find-by-name:%v"
	// // RegencyFindAllTotal is a redis key for find all regency.
	// RegencyFindAllTotal = prefix + ":regency:find-all-total:%v:%v:%v:%v:%v:%v:%v"
	// // RegencyFindAll is a redis key for find all regency.
	// RegencyFindAll = prefix + ":regency:find-all:%v:%v:%v:%v:%v:%v:%v"
	// // RegencyFindByProvinceID is a redis key for find regency by province id.
	// RegencyFindByProvinceID = prefix + ":regency:find-by-province-id:%v:%v"
	// // VillageFindByDistrictID is a redis key for find village by district id.
	// VillageFindByDistrictID = prefix + ":village:find-by-district-id:%v:%v"
	// // PageFindBySlug is a redis key for find page by slug.
	// PageFindBySlug = prefix + ":page:find-by-slug:%v"
	// // PaymentChannelFindAll is a redis key for find all payment channel.
	// PaymentChannelFindAll = prefix + ":payment-channel:find-all"
	// UserFindByEmail is a redis key for find user by email.
	// UserFindByEmail = prefix + ":user:find-by-email:%v"

	// UserExistsByEmail is a redis key for exists user by email.
	// UserExistsByEmail = prefix + ":user:exists-by-email:%v"
	// UserFindByUSername is a redis key for find user by username.
	UserFindByUSername = prefix + ":user:find-by-username:%v"
	// // UserEmployeeFindByMerchantID is a redis key for find user employee by merchant id.
	// UserEmployeeFindByMerchantID = prefix + ":user:find-by-merchant-id:%v:%v:%v:%v:%v:%v:%v"
	// // UserEmployeeFindByMerchantIDTotal is a redis key for find user employee by merchant id total.
	// UserEmployeeFindByMerchantIDTotal = prefix + ":user:find-by-merchant-id-total:%v:%v:%v:%v:%v:%v:%v"
	// // FaqFindAll is a redis key for find all faq.
	// FaqFindAll = prefix + ":faq:find-all:%v:%v:%v:%v:%v:%v:%v:%v"
	// // FaqFindAllTotal is a redis key for find all faq total.
	// FaqFindAllTotal = prefix + ":faq:find-all-total:%v:%v:%v:%v:%v:%v:%v:%v"
	// // FaqFindByID is a redis key for find faq by id.
	// FaqFindByID = prefix + ":faq:find-by-id:%v"
	// // FaqCategoryFindAllActive is a redis key for find all faq category active.
	// FaqCategoryFindAllActive = prefix + ":faq-category:find-all-active"
	// // RegencyFindByID is a redis key for find regency by id.
	// RegencyFindByID = prefix + ":regency:find-by-id:%v"
	// // RegencyFindByName is a redis key for find regency by name.
	// RegencyFindByName = prefix + ":regency:find-by-name:%v"
	// // DistrictFindByID is a redis key for find district by id.
	// DistrictFindByID = prefix + ":district:find-by-id:%v"
	// // VillageFindByID is a redis key for find village by id.
	// VillageFindByID = prefix + ":village:find-by-id:%v"
	// // RoleOwnerFindAll is a redis key for find all role owner.
	// RoleOwnerFindAll = prefix + ":role-owner:find-all"
	// // UserOwnerFindByMerchantID is a redis key for find user owner by merchant id.
	// UserOwnerFindByMerchantID = prefix + ":user-owner:find-by-merchant-id:%v:%v"
	// // MerchantFindByID is a redis key for find merchant by id.
	// MerchantFindByID = prefix + ":merchant:find-by-id:%v"
	// // InvoiceFindByTrxID is a redis key for find invoice by trx id.
	// InvoiceFindByTrxID = prefix + ":invoice:find-by-trx-id:%v"
	// UserFindByMerchantID is a redis key for find user by merchant id.
	// UserFindByMerchantID = prefix + ":user:find-by-merchant-id:%v"
	// UserRoleFindByUserIDRoleID is a redis key for find user role by user id and role id.
	UserRoleFindByUserIDRoleID = prefix + ":user-role:find-by-user-id:%v:%v"
	// TokenUserByJTI is a redis key for find token user by jti.
	TokenUserByJTI = prefix + ":token-user:find-by-jti2:%v"

	// // ReminderChangePIN is a redis key for find reminder change pin.
	// ReminderChangePIN = prefix + ":reminder:change-pin:%v"
	// // UserDetailFindByUserID is a redis key for find user detail by user id.
	// UserDetailFindByUserID = prefix + ":user-detail:find-by-user-id:%v"
	// ProductFindByID is a redis key for find product by id.
	ProductFindByID = prefix + ":product:find-by-id:%v"
	// // CartFindAll is a redis key for find all cart.
	// CartFindAll = prefix + ":cart:find-all:%v"
	// // MerchantFindByIDs is a redis key for find merchant by ids.
	// MerchantFindByIDs = prefix + ":merchant:find-by-ids:%v"
	// // UserAddressFindByID is a redis key for find user address by id.
	// UserAddressFindByID = prefix + ":user-address:find-by-id:%v"
	// // UserAddressFindByUserID is a redis key for find user address by user id.
	// UserAddressFindByUserID = prefix + ":user-address:find-by-user-id:%v"
	// // UserAddressPickupPointByUserID is a redis key for find user address pickup point.
	// UserAddressPickupPointByUserID = prefix + ":user-address:pickup-point:%v"
	// // CartFindByProductIDAndUserID is a redis key for find cart by product id and user id.
	// CartFindByProductIDAndUserID = prefix + ":cart:find-by-product-id-and-user-id:%v:%v"
	// // OrderFindByID is a redis key for find order by id.
	// OrderFindByID = prefix + ":order:find-by-id:%v"
	// // PaymentChannelFindByID is a redis key for find payment channel by id.
	// PaymentChannelFindByID = prefix + ":payment-channel:find-by-id:%v"
	// // MerchantProductFindByMerchantAndProductIDs is a redis key for find merchant product by merchant and product ids.
	// MerchantProductFindByMerchantAndProductIDs = prefix + ":merchant-product:find-by-merchant-and-product-ids:%v:%v"
	// // CartFindByIDs is a redis key for find cart by ids.
	// CartFindByIDs = prefix + ":cart:find-by-ids:%v"
	// // CartFindByID is a redis key for find cart by id.
	// CartFindByID = prefix + ":cart:find-by-id:%v"
	// // MerchantFindByAddressTotal is a redis key for find merchant by address total.
	// MerchantFindByAddressTotal = prefix + ":merchant:find-by-address:%v"
	// // MerchantFindByAddress is a redis key for find merchant by address.
	// MerchantFindByAddress = prefix + ":merchant:find-by-address-total:%v"
	// // MerchantFindByNameTotal is a redis key for find merchants by name total.
	// MerchantFindByNameTotal = prefix + ":merchant:find-by-name:%v"
	// // MerchantFindByName is a redis key for find merchants by name.
	// MerchantFindByName = prefix + ":merchant:find-by-name-total:%v"
	// // OrderFindAllByUserIDTotal is a redis key for find all order by user id total.
	// OrderFindAllByUserIDTotal = prefix + ":order:find-all-by-user-id-total:%v:%v:%v:%v:%v:%v:%v:%v:%v"
	// // OrderFindAllByUserID is a redis key for find all order by user id.
	// OrderFindAllByUserID = prefix + ":order:find-all-by-user-id:%v:%v:%v:%v:%v:%v:%v:%v:%v"
	// // PertaminaProductFindByID is a redis key for find pertamina product by id.
	// PertaminaProductFindByID = prefix + ":pertamina-product:find-by-id:%v"
	// // PertaminaProductFindByNameAndShortDescription is a redis key for find pertamina product by name and short description.
	// PertaminaProductFindByNameAndShortDescription = prefix + ":pertamina-product:find-by-name-and-short-description:%v:%v"
	// // PertaminaProductFindByName is a redis key for find pertamina product by name.
	// PertaminaProductFindByName = prefix + ":pertamina-product:find-by-name:%v"
	// // PertaminaProductFindByExternalID is a redis key for find pertamina product by external id.
	// PertaminaProductFindByExternalID = prefix + ":pertamina-product:find-by-external-id:%v"
	// // MerchantProductFindLatestUpdatedByMerchantID is a redis key for find merchant product latest updated by merchant id.
	// MerchantProductFindLatestUpdatedByMerchantID = prefix + ":merchant-product:find-latest-updated-by-merchant-id:%v"
	// // CustomerFindLatestUpdatedByMerchantID is a redis key for find customer latest updated by merchant id.
	// CustomerFindLatestUpdatedByMerchantID = prefix + ":customer:find-latest-updated-by-merchant-id:%v"
	// UserRoleFindByRoleID is a redis key for find user role by role id.
	UserRoleFindByRoleID = prefix + ":user-role:find-by-role-id:%v"
	// // 	UserFindByUserIDs is a redis key for find user by user ids.
	// UserFindByUserIDs = prefix + ":user:find-by-user-ids:%v"
	// // UserFindLatestUpdateByMerchantID is a redis key for find user latest update by merchant id.
	// UserFindLatestUpdateByMerchantID = prefix + ":user:find-latest-update-by-merchant-id:%v"
	// // OrderFindTransactionOrderCMSTotal is a redis key for find transaction order cms total.
	// OrderFindTransactionOrderCMSTotal = prefix + ":order:find-transaction-order-cms-total:%v:%v:%v"
	// RoleFindByType is a redis key for find role by type.
	RoleFindByType = prefix + ":role:find-by-type:%v:%v"

// // RegionFindByID is a redis key for find region by id.
// RegionFindByID = prefix + ":region:find-by-id:%v"
// // RegionFindByName is a redis key for find region by name.
// RegionFindByName = prefix + ":region:find-by-name:%v"
// // MerchantFindAll is a redis key for find all merchant.
// MerchantFindAll = prefix + ":merchant:find-all:%v:%v:%v:%v:%v"
// // MerchantFindAllCMS is a redis key for find all merchant cms.
// MerchantFindAllCMS = prefix + ":merchant:find-all-cms:%v:%v:%v"
// // MerchantFindAllTotalCMS is a redis key for find all merchant total.
// MerchantFindAllTotalCMS = prefix + ":merchant:find-all-total:%v"
// // MenuFindWithPermission is a redis key for find menu with permission.
// MenuFindWithPermission = prefix + ":menu:find-with-permission:%v"
// // MerchantFindAllTotal is a redis key for find all merchant total.
// MerchantFindAllTotal = prefix + ":merchant:find-all-total:%v:%v:%v:%v:%v"
// // PaymentChannelFindByCode is a redis key for find payment channel by code.
// PaymentChannelFindByCode = prefix + ":payment-channel:find-by-code:%v"
// // SettingFindBySlug
// SettingFindBySlug = prefix + ":setting:find-by-slug:%v"
)
