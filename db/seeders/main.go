package main

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/utils"
)

var adminRoleID uuid.UUID
var employeeRoleID uuid.UUID
var shiftID uuid.UUID

func main() {
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	db, err := utils.NewPostgresGormDB(cfg, nil)
	checkError(err)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// createRole(db)
	// createAdmin(db)
	// createEmployee(db, 5)
	createShift(db)
	createShiftDetail(db)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createRole(db *gorm.DB) {
	adminRoleID = uuid.New()
	employeeRoleID = uuid.New()
	roles := []*entity.Role{
		entity.NewRole(employeeRoleID, "employee", "system"),
		entity.NewRole(adminRoleID, "admin", "system"),
	}

	if err := db.Model(&entity.Role{}).Create(roles).Error; err != nil {
		panic(err)
	}
}

func createAdmin(db *gorm.DB) {
	admin := entity.NewUser(
		uuid.New(),
		"admin",
		"admin123",
		0,
		adminRoleID,
		"system",
	)

	if err := db.Model(&entity.User{}).
		Create(admin).
		Error; err != nil {
		panic(err)
	}
}

func createEmployee(db *gorm.DB, total int) {
	min := 2000000
	max := 20000000
	employees := make([]*entity.User, 0)
	for i := 0; i < total; i++ {
		employee := entity.NewUser(
			uuid.New(),
			"employee"+strconv.Itoa(i),
			"employee123",
			rand.Intn(max-min+1)+min,
			employeeRoleID,
			"system",
		)
		employees = append(employees, employee)
	}

	if err := db.Model(&entity.User{}).
		Create(employees).
		Error; err != nil {
		panic(err)
	}
}

func createShift(db *gorm.DB) {
	shiftID = uuid.New()
	shift := entity.NewShift(
		shiftID,
		"Morning Shift",
		"system",
	)

	if err := db.Model(&entity.Shift{}).
		Create(shift).
		Error; err != nil {
		panic(err)
	}
}

func createShiftDetail(db *gorm.DB) {
	shiftDetails := make([]*entity.ShiftDetail, 0)
	for i := 1; i <= 7; i++ {
		if i >= 1 && i <= 5 {
			shiftDetail := entity.NewShiftDetail(
				uuid.New(),
				shiftID,
				"ON",
				true,
				i,
				"09:00:00",
				"17:00:00",
				"system",
			)
			shiftDetails = append(shiftDetails, shiftDetail)
		} else {
			shiftDetail := entity.NewShiftDetail(
				uuid.New(),
				shiftID,
				"OFF",
				false,
				i,
				"00:00:00",
				"00:00:00",
				"system",
			)
			shiftDetails = append(shiftDetails, shiftDetail)
		}
	}

	if err := db.Model(&entity.ShiftDetail{}).
		Create(shiftDetails).
		Error; err != nil {
		panic(err)
	}
}
