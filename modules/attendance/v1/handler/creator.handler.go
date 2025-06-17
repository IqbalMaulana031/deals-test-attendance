package handler

import (
	"net/http"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/errors"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/attendance/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

// AttendanceCreatorHandler is a handler for attendance finder
type AttendanceCreatorHandler struct {
	cfg               config.Config
	attendanceCreator service.AttendanceCreatorUseCase
	cache             interfaces.Cacheable
}

// NewAttendanceCreatorHandler is a constructor for AttendanceCreatorHandler
func NewAttendanceCreatorHandler(
	cfg config.Config,
	attendanceCreator service.AttendanceCreatorUseCase,
	cache interfaces.Cacheable,
) *AttendanceCreatorHandler {
	return &AttendanceCreatorHandler{
		cfg:               cfg,
		attendanceCreator: attendanceCreator,
		cache:             cache,
	}
}

// CreateAttendance is a handler for create attendance
func (ach *AttendanceCreatorHandler) RechordTime(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "RechordTime")
	defer span.End()
	requestId := c.GetString("requestId")

	var attendance resource.AttendanceRequest
	if err := c.ShouldBindJSON(&attendance); err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	// Create attendance
	err := ach.attendanceCreator.RechordTime(ctx, uuid.MustParse(c.Request.Header.Get(constant.RequestHeaderUserID)), attendance.ShiftDetailID)
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", nil))
}
