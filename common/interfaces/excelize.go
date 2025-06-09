package interfaces

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// ExcelizeUseCase define interface for excelize
type ExcelizeUseCase interface {
	WriteExcelFromStruct(ctx *gin.Context, sheetName string, data interface{}) error
	WriteExcelFromMap(ctx *gin.Context, sheetName string, data []map[string]interface{}, headers []string) error
	ReadExcel(ctx *gin.Context, file multipart.File) ([][]string, error)
	WriteExcelMultipleSheet(ctx *gin.Context, sheetName []string, data []interface{}, headers []string) error
}
