package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/utils"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	db, err := utils.NewPostgresGormDB(cfg, nil)
	checkError(err)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	createProductCategoryAndPertaminaProduct(db)
	createRegion(db)
	createBusinessField(db)
	createTypeMerchant(db)
	createBusinessType(db)
	createPermission(db)
	createSuperAdmin(db)
	createFaqCategory(db)
	// The code below is commented out because synchronization with Firestore is no longer needed
	// createMerchantLocationFirestore(cfg, db)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createFaqCategory(db *gorm.DB) {
	faqCategory := entity.NewFaqCategory(
		uuid.New(),
		"Transaksi",
		"transaksi",
		true,
		"system",
	)

	if err := db.Model(&entity.FaqCategory{}).Create(faqCategory).Error; err != nil {
		panic(err)
	}
}

func createSuperAdmin(db *gorm.DB) {
	superAdmin := entity.NewUser(
		uuid.New(),
		utils.StringToNullString(uuid.New().String()),
		"superadmin",
		"superadmin@deals.id",
		"superadmin",
		"",
		"0",
		"bandung",
		utils.TimeToNullTime(time.Time{}),
		utils.StringToNullString(""),
		utils.StringToNullString(""),
		"system",
	)

	if err := db.Model(&entity.User{}).
		Create(superAdmin).
		Error; err != nil {
		panic(err)
	}
}

func createProductCategoryAndPertaminaProduct(db *gorm.DB) {
	type productCategory struct {
		Name      string
		Type      string
		StockType string
	}

	category := []*productCategory{
		{Name: "LPG", Type: "pertamina", StockType: "tabung"},
		{Name: "BBM", Type: "pertamina", StockType: "liter"},
		{Name: "Pelumas", Type: "pertamina", StockType: "botol"},
		{Name: "Refrigeran", Type: "pertamina", StockType: "kg"},
		{Name: "Non Pertamina", Type: "non_pertamina", StockType: "unit"},
	}

	productCategories := make([]*entity.ProductCategory, 0)
	for _, c := range category {
		productCategories = append(productCategories, entity.NewProductCategory(
			uuid.New(),
			c.Name,
			c.Type,
			c.StockType,
			"system",
		))
	}

	if err := db.WithContext(context.Background()).
		Model(&entity.ProductCategory{}).
		Create(productCategories).
		Error; err != nil {
		panic(err)
	}

	createPertaminaProduct(db)
}

func createPertaminaProduct(db *gorm.DB) {
	type product struct {
		Name      string
		StockType string
	}

	type category struct {
		Name     string
		products []product
	}

	categories := make([]category, 0)
	lpg := []product{
		{Name: "Bright Gas 5,5 Kg", StockType: "tabung"},
		{Name: "Bright Gas 12 Kg", StockType: "tabung"},
		{Name: "LPG 3 Kg", StockType: "tabung"}}
	categories = append(categories, category{Name: "LPG", products: lpg})

	bbm := []product{
		{Name: "Pertalite", StockType: "liter"},
		{Name: "Pertamax - Jerrycan", StockType: "liter"},
		{Name: "Bio Solar", StockType: "liter"},
		{Name: "Pertamax", StockType: "liter"},
		{Name: "DEXLITE", StockType: "liter"},
		{Name: "Pertamina DEX", StockType: "liter"},
		{Name: "Pertamina Dex - Jerrycan", StockType: "liter"},
		{Name: "Dexlite - Jerrycan", StockType: "liter"},
		{Name: "Pertamax Turbo - Jerrycan", StockType: "liter"},
		{Name: "Pertamax Turbo", StockType: "liter"},
	}
	categories = append(categories, category{Name: "BBM", products: bbm})

	pelumas := []product{
		{Name: "Dr. Lube", StockType: "botol"},
		{Name: "Fastron", StockType: "botol"},
		{Name: "Enduro Racing", StockType: "botol"},
	}
	categories = append(categories, category{Name: "Pelumas", products: pelumas})

	for _, c := range categories {
		category := &entity.ProductCategory{}
		if err := db.Model(&entity.ProductCategory{}).
			Where("name = ?", c.Name).
			First(category).Error; err != nil {
			panic(err)
		}

		for _, p := range c.products {
			product := entity.NewPertaminaProduct(
				uuid.New(),
				category.ID,
				p.Name,
				"",
				p.StockType,
				"system",
			)

			if err := db.Model(&entity.PertaminaProduct{}).
				Create(product).Error; err != nil {
				panic(err)
			}
		}
	}
}

func createRegion(db *gorm.DB) {
	regions := []string{
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
	}

	regionEntities := make([]*entity.Region, 0)
	for _, r := range regions {
		regionEntities = append(regionEntities, entity.NewRegion(
			uuid.New(),
			r,
			strings.ToLower(r),
			"system",
		))
	}

	if err := db.Model(&entity.Region{}).
		Create(regionEntities).
		Error; err != nil {
		panic(err)
	}
}

func createBusinessField(db *gorm.DB) {
	list := []string{
		"Dealer Mobil dan atau Dealer Motor",
		"Event Organizer",
		"Fashion dan Kecantikan",
		"Jasa",
		"Kantor Layanan Pemerintah",
		"Konstruksi",
		"Kuliner",
		"Lainnya",
		"Makanan & Minuman",
		"Minyak Bumi dan Produk Minyak Bumi",
		"Organisasi Keagamaan",
		"Otomotif",
		"Pariwisata",
		"Percetakan",
		"Perguruan Tinggi, Universitas, Sekolah Profesi & Sekolah Menengah Pertama",
		"Pertambangan",
		"Pertanian dan Perkebunan",
		"Peternakan dan Perikanan",
		"Salon & Tempat Pangkas Rambut",
		"Stasiun Layanan",
		"Stasiun Pengisian Bahar Bakar Minyak & Gas",
		"Teknologi, Informasi dan Komunikasi",
		"Tempat Parkir Mobil dan Garasi",
		"Toko Gadget & Elektronik",
		"Toko Kelontong atau Supermarket",
		"Toko Kosmetik",
		"Toko Peralatan Rumah Tangga",
		"Tour & Travel",
		"Transportasi",
	}

	fields := make([]*entity.BusinessField, 0)
	for _, l := range list {
		fields = append(fields, entity.NewBusinessField(
			uuid.New(),
			l,
			slug.Make(l),
			"system",
		))
	}
	if err := db.Model(&entity.BusinessField{}).Create(fields).
		Error; err != nil {
		panic(err)
	}
}

//nolint:funlen
func createPermission(db *gorm.DB) {
	can := "can"
	listCRUD := []string{"create", "read", "update", "delete", "activate", "cancel", "reset", "download", "share"}
	listMenu := []string{"auth", "account", "product", "order", "report", "merchant_info"}
	listMenuLabel := []string{"Autentikasi", "Akun Usaha", "Manage Produk", "Transaksi", "Laporan", "Info Merchant"}

	// create menu
	menus := make([]*entity.Menu, 0)
	for i, m := range listMenu {
		menus = append(menus, entity.NewMenu(
			i+1,
			m,
			listMenuLabel[i],
			"system",
		))
	}

	if err := db.Model(&entity.Menu{}).Create(menus).
		Error; err != nil {
		panic(err)
	}

	// create permission for menu
	permissions := make([]*entity.Permission, 0)

	listPermissionAuth := []string{
		can + "_" + "login",
		can + "_" + "logout",
	}

	for _, l := range listPermissionAuth {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[0].ID,
			"system",
		))
	}

	listPermissionAccount := []string{
		can + "_" + listCRUD[0] + "_employee",
		can + "_" + listCRUD[1] + "_employee",
		can + "_" + listCRUD[2] + "_employee",
		can + "_" + listCRUD[3] + "_employee",
		can + "_" + listCRUD[0] + "_customer",
		can + "_" + listCRUD[1] + "_customer",
		can + "_" + listCRUD[2] + "_customer",
		can + "_" + listCRUD[3] + "_customer",
		can + "_" + listCRUD[2] + "_profile",
		can + "_" + listCRUD[6] + "_pin",
		can + "_" + listCRUD[1] + "_permission_employee",
	}

	for _, l := range listPermissionAccount {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[1].ID,
			"system",
		))
	}

	listPermissionProduct := []string{
		can + "_" + listCRUD[0] + "_product",
		can + "_" + listCRUD[1] + "_product",
		can + "_" + listCRUD[2] + "_product",
		can + "_" + listCRUD[3] + "_product",
		can + "_" + listCRUD[4] + "_category_product",
	}

	for _, l := range listPermissionProduct {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[2].ID,
			"system",
		))
	}

	listPermissionOrder := []string{
		can + "_" + listCRUD[0] + "_order",
		can + "_" + listCRUD[5] + "_order",
	}

	for _, l := range listPermissionOrder {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[3].ID,
			"system",
		))
	}

	listPermissionReport := []string{
		can + "_" + listCRUD[1] + "_report_profit",
		can + "_" + listCRUD[1] + "_report_best_product",
		can + "_" + listCRUD[1] + "_report_customer_sales",
		can + "_" + listCRUD[1] + "_report_operational_cost",
		can + "_" + listCRUD[7] + "_report",
		can + "_" + listCRUD[8] + "_report",
	}

	for _, l := range listPermissionReport {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[4].ID,
			"system",
		))
	}

	listPermissionMerchantInfo := []string{
		can + "_" + listCRUD[2] + "_merchant_info",
		can + "_" + listCRUD[2] + "_operation_time",
	}

	for _, l := range listPermissionMerchantInfo {
		permissions = append(permissions, entity.NewPermission(
			uuid.New(),
			l,
			slug.Make(l),
			//nolint
			menus[5].ID,
			"system",
		))
	}

	if err := db.Model(&entity.Permission{}).Create(permissions).
		Error; err != nil {
		panic(err)
	}

	// find all role
	roles := make([]*entity.Role, 0)
	if err := db.WithContext(context.Background()).Model(&entity.Role{}).Find(&roles).Error; err != nil {
		log.Println("error get role")
		panic(err)
	}

	listMenuOwner := []string{
		can + "_" + "login",
		can + "_" + "logout",
		can + "_" + listCRUD[0] + "_product",
		can + "_" + listCRUD[1] + "_product",
		can + "_" + listCRUD[2] + "_product",
		can + "_" + listCRUD[3] + "_product",
		can + "_" + listCRUD[4] + "_category_product",
		can + "_" + listCRUD[0] + "_order",
		can + "_" + listCRUD[5] + "_order",
		can + "_" + listCRUD[1] + "_report_profit",
		can + "_" + listCRUD[1] + "_report_best_product",
		can + "_" + listCRUD[1] + "_report_customer_sales",
		can + "_" + listCRUD[1] + "_report_operational_cost",
		can + "_" + listCRUD[0] + "_employee",
		can + "_" + listCRUD[1] + "_employee",
		can + "_" + listCRUD[2] + "_employee",
		can + "_" + listCRUD[3] + "_employee",
		can + "_" + listCRUD[0] + "_customer",
		can + "_" + listCRUD[1] + "_customer",
		can + "_" + listCRUD[2] + "_customer",
		can + "_" + listCRUD[3] + "_customer",
		can + "_" + listCRUD[2] + "_profile",
		can + "_" + listCRUD[6] + "_pin",
		can + "_" + listCRUD[1] + "_permission_employee",
		can + "_" + listCRUD[7] + "_report",
		can + "_" + listCRUD[8] + "_report",
		can + "_" + listCRUD[2] + "_merchant_info",
		can + "_" + listCRUD[2] + "_operation_time",
	}

	listMenuEmployeeAdmin := []string{
		can + "_" + "login",
		can + "_" + "logout",
		can + "_" + listCRUD[0] + "_product",
		can + "_" + listCRUD[1] + "_product",
		can + "_" + listCRUD[2] + "_product",
		can + "_" + listCRUD[3] + "_product",
		can + "_" + listCRUD[4] + "_category_product",
		can + "_" + listCRUD[0] + "_order",
		can + "_" + listCRUD[5] + "_order",
		can + "_" + listCRUD[1] + "_report_profit",
		can + "_" + listCRUD[1] + "_report_best_product",
		can + "_" + listCRUD[1] + "_report_customer_sales",
		can + "_" + listCRUD[0] + "_customer",
		can + "_" + listCRUD[1] + "_customer",
		can + "_" + listCRUD[2] + "_customer",
		can + "_" + listCRUD[3] + "_customer",
		can + "_" + listCRUD[2] + "_merchant_info",
		can + "_" + listCRUD[2] + "_operation_time",
	}

	listMenuEmployeeCashier := []string{
		can + "_" + "login",
		can + "_" + "logout",
		can + "_" + listCRUD[0] + "_order",
		can + "_" + listCRUD[5] + "_order",
		can + "_" + listCRUD[1] + "_report_profit",
		can + "_" + listCRUD[1] + "_report_best_product",
		can + "_" + listCRUD[1] + "_report_customer_sales",
		can + "_" + listCRUD[0] + "_customer",
		can + "_" + listCRUD[1] + "_customer",
		can + "_" + listCRUD[2] + "_customer",
		can + "_" + listCRUD[3] + "_customer",
	}

	// create role permission
	for _, role := range roles {
		rolePermission := make([]*entity.RolePermission, 0)
		switch role.Name {
		case "owner":
			for _, l := range listMenuOwner {
				permissionEnt := &entity.Permission{}
				if err := db.WithContext(context.Background()).Model(&entity.Permission{}).Where("name = ?", l).First(permissionEnt).Error; err != nil {
					log.Println("error get permission")
					panic(err)
				}

				rolePermission = append(rolePermission, entity.NewRolePermission(
					uuid.New(),
					role.ID,
					permissionEnt.ID,
					"system",
				))
			}
		case "employee_admin":
			for _, l := range listMenuEmployeeAdmin {
				permissionEnt := &entity.Permission{}
				if err := db.WithContext(context.Background()).Model(&entity.Permission{}).Where("name = ?", l).First(permissionEnt).Error; err != nil {
					log.Println("error get permission")
					panic(err)
				}
				rolePermission = append(rolePermission, entity.NewRolePermission(
					uuid.New(),
					role.ID,
					permissionEnt.ID,
					"system",
				))
			}

		case "employee_cashier":
			for _, l := range listMenuEmployeeCashier {
				permissionEnt := &entity.Permission{}
				if err := db.WithContext(context.Background()).Model(&entity.Permission{}).Where("name = ?", l).First(permissionEnt).Error; err != nil {
					log.Println("error get permission")
					panic(err)
				}
				rolePermission = append(rolePermission, entity.NewRolePermission(
					uuid.New(),
					role.ID,
					permissionEnt.ID,
					"system",
				))
			}
		}

		if len(rolePermission) > 0 {
			if err := db.Model(&entity.RolePermission{}).Create(rolePermission).
				Error; err != nil {
				log.Println("error")
				panic(err)
			}
		}
	}
}

// TODO: Comment this code because unnecessary set merchant to firestore
// func createMerchantLocationFirestore(cfg *config.Config, db *gorm.DB) {
// 	ctx := context.Background()
// 	fs := firestore.NewFirestore(*cfg)
// 	merchants := make([]*entity.Merchant, 0)
// 	if err := db.WithContext(ctx).
// 		Model(&entity.Merchant{}).
// 		Find(&merchants).Error; err != nil {
// 		log.Println("error get merchant")
// 		panic(err)
// 	}

// 	for _, m := range merchants {
// 		loc := &map[string]interface{}{
// 			"latitude":  m.Latitude,
// 			"longitude": m.Longitude,
// 			"geohash":   utils.GeohashForLocation(m.Latitude, m.Longitude, 0),
// 		}

// 		if err := fs.Set(ctx,
// 			constant.MerchantLocationFirestoreCollection,
// 			m.ID.String(),
// 			loc,
// 		); err != nil {
// 			log.Println("error set merchant location to firestore", err)
// 			panic(err)
// 		}
// 	}
// }

// createTypeMerchant create type merchant
func createTypeMerchant(db *gorm.DB) {
	list := []string{
		"Agen Minyak Tanah",
		"Bengkel",
		"Bright Cafe",
		"Bright Store",
		"Kios Pertamina",
		"Lainnya",
		"Pangkalan LPG",
		"Pertashop",
		"Pertashop COCO",
		"SPBU",
		"SPBU Nelayan",
		"Toko/ Merchant",
		"UMKM Mitra binaan",
	}

	fields := make([]*entity.TypeMerchant, 0)
	for _, l := range list {
		fields = append(fields, entity.NewTypeMerchant(
			uuid.New(),
			l,
			true,
			slug.Make(l),
			"system",
		))
	}
	if err := db.Model(&entity.TypeMerchant{}).Create(fields).
		Error; err != nil {
		panic(err)
	}
}

// createBusinessType create business type
func createBusinessType(db *gorm.DB) {
	list := []string{
		"CV",
		"Firma",
		"Koperasi",
		"Lainnya",
		"PT",
		"Perorangan",
		"Yayasan",
	}

	fields := make([]*entity.BusinessType, 0)
	for _, l := range list {
		fields = append(fields, entity.NewBusinessType(
			uuid.New(),
			l,
			slug.Make(l),
			"system",
		))
	}
	if err := db.Model(&entity.BusinessType{}).Create(fields).
		Error; err != nil {
		panic(err)
	}
}
