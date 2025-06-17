package handler

import (
	"fmt"
	"net/http"
	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/middleware"
	"starter-go-gin/modules/attendance/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

// AttendanceFinderHandler is a handler for product finder
type AttendanceFinderHandler struct {
	attendanceFinder service.AttendanceFinderUseCase
	cfg              config.Config
}

// NewAttendanceFinderHandler is a constructor for AttendanceFinderHandler
func NewAttendanceFinderHandler(
	attendanceFinder service.AttendanceFinderUseCase,
	cfg config.Config,
) *AttendanceFinderHandler {
	return &AttendanceFinderHandler{
		attendanceFinder: attendanceFinder,
		cfg:              cfg,
	}
}

// GetAttendanceByID is a handler for get attendance by ID
func (afh *AttendanceFinderHandler) GetAttendanceByID(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "GetAttendanceByID")
	defer span.End()
	requestId := c.GetString("requestId")

	attendanceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	attendance, err := afh.attendanceFinder.GetAttendanceByID(ctx, attendanceID)
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", attendance))
}

// GetAttendance is a handler for get attendance
func (afh *AttendanceFinderHandler) GetAttendance(c *gin.Context) {
	fmt.Println("GetShift called")
	ctx, span := trace.StartSpan(c.Request.Context(), "GetShift")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	var request resource.PaginationQueryParam

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		c.Abort()
		return
	}

	// Get shifts by filter
	Attendance, total, err := afh.attendanceFinder.GetAttendance(ctx, request.Query, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	totalPage := int(total) / request.Limit
	if int(total)%request.Limit > 0 {
		totalPage++
	}

	res := &resource.GetAttendanceListResponse{
		List:  Attendance,
		Total: total,
		Meta: &resource.Meta{
			Total:       int(total),
			Limit:       request.Limit,
			Page:        request.Page,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
		},
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", res))
}

// GetAttendanceByUserID is a handler for get attendance by user ID
func (afh *AttendanceFinderHandler) GetAttendanceByUserID(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "GetAttendanceByUserID")
	defer span.End()
	requestId := c.GetString("requestId")

	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	var request resource.PaginationQueryParam

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		c.Abort()
		return
	}

	// Get attendance by user ID
	attendance, total, err := afh.attendanceFinder.GetAttendanceByUserID(ctx, userID, request.Query, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	totalPage := int(total) / request.Limit
	if int(total)%request.Limit > 0 {
		totalPage++
	}

	res := &resource.GetAttendanceListResponse{
		List:  attendance,
		Total: total,
		Meta: &resource.Meta{
			Total:       int(total),
			Limit:       request.Limit,
			Page:        request.Page,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
		},
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", res))
}
