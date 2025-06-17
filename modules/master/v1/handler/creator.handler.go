package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/master/v1/service"
)

// MasterCreatorHandler is a handler for master finder
type MasterCreatorHandler struct {
	cfg           config.Config
	masterCreator service.MasterCreatorUseCase
	cache         interfaces.Cacheable
}

// NewMasterCreatorHandler is a constructor for MasterCreatorHandler
func NewMasterCreatorHandler(
	cfg config.Config,
	masterCreator service.MasterCreatorUseCase,
	cache interfaces.Cacheable,
) *MasterCreatorHandler {
	return &MasterCreatorHandler{
		cfg:           cfg,
		masterCreator: masterCreator,
		cache:         cache,
	}
}

// // CreateMaster is a handler for create master
// func (pch *MasterCreatorHandler) CreateMaster(c *gin.Context) {
// 	ctx, span := trace.StartSpan(c.Request.Context(), "handler.masterCreator.Create")
// 	defer span.End()

// 	var req resource.CreateMasterRequest

// 	if err := c.ShouldBind(&req); err != nil {
// 		parseError := errors.ParseErrorValidation(err)
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
// 		c.Abort()
// 		return
// 	}

// 	master := entity.NewMaster(
// 		uuid.New(),
// 		req.MasterName,
// 		req.Stock,
// 		req.Price,
// 		"system",
// 	)
// 	err := pch.masterCreator.CreateMaster(ctx, master)

// 	if err != nil {
// 		parseError := errors.ParseError(err)
// 		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
// }
