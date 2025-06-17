package handler

import (
	"fmt"
	"net/http"
	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/middleware"
	"starter-go-gin/modules/master/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

// MasterFinderHandler is a handler for product finder
type MasterFinderHandler struct {
	masterFinder service.MasterFinderUseCase
	cfg          config.Config
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func NewMasterFinderHandler(
	masterFinder service.MasterFinderUseCase,
	cfg config.Config,
) *MasterFinderHandler {
	return &MasterFinderHandler{
		masterFinder: masterFinder,
		cfg:          cfg,
	}
}

func (mfh *MasterFinderHandler) GetShiftByID(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "GetShiftByID")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	id := c.Param("id")

	// Get shift by ID
	shift, err := mfh.masterFinder.GetShiftByID(ctx, uuid.MustParse(id))
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	res := resource.Shift{
		ID:        shift.ID,
		ShiftName: shift.ShiftName,
		IsDefault: shift.IsDefault,
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", res))
}

func (mfh *MasterFinderHandler) GetShift(c *gin.Context) {
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
	shifts, total, err := mfh.masterFinder.GetShift(ctx, request.Query, request.Sort, request.Order, request.Limit, request.Page)
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

	res := &resource.GetShiftListResponse{
		List:  shifts,
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

func (mfh *MasterFinderHandler) GetShiftAndDetailsByID(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "GetShiftAndDetailsByID")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	id := c.Param("id")

	// Get shift and details by ID
	shift, err := mfh.masterFinder.GetShiftAndDetailsByID(ctx, uuid.MustParse(id))
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	res := resource.GetShiftAndDetailsByIDResponse{
		ID:        shift.ID,
		ShiftName: shift.ShiftName,
		IsDefault: shift.IsDefault,
		Details:   shift.ShiftDetails,
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", res))
}

// GetShiftDetailsByID is a handler for getting shift details by ID
func (mfh *MasterFinderHandler) GetShiftDetailsByID(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "GetShiftDetailsByID")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	var request resource.GetShiftByIDRequest

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		c.Abort()
		return
	}

	// Get shift details by ID
	shift, err := mfh.masterFinder.GetShiftDetailByID(ctx, request.ID)
	if err != nil {
		parser := errors.ParseError(err)
		c.JSON(parser.Code, response.ErrorAPIResponse(parser.Code, requestId, parser.Message))
		c.Abort()
		return
	}

	res := resource.ShiftDetail{
		ID:          shift.ID,
		ShiftID:     shift.ShiftID,
		Code:        shift.Code,
		DayType:     shift.DayType,
		DayInNumber: shift.DayInNumber,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", res))
}
