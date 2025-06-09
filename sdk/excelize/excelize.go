package excelize

import (
	"fmt"
	"mime/multipart"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"

	"starter-go-gin/common/constant"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
)

var character []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// Excelize is an struct for Excelize SDK
type Excelize struct {
	cfg config.Config
}

// NewExcelize initiate Excelize SDK
func NewExcelize(cfg config.Config) *Excelize {
	return &Excelize{
		cfg: cfg,
	}
}

// WriteExcelFromStruct is a function to write data to excel file from struct
func (e *Excelize) WriteExcelFromStruct(ctx *gin.Context, sheetName string, data interface{}) error {
	reflectValue := reflect.ValueOf(data)

	ret := make([]interface{}, reflectValue.Len())

	for i := 0; i < reflectValue.Len(); i++ {
		ret[i] = reflectValue.Index(i).Interface()
	}

	if len(ret) == 0 {
		return errors.Wrap(errors.New("Data is empty"), "[Excelize-WriteExcel] error while write excel")
	}

	sampleReflectValue := reflect.ValueOf(ret[0])
	if sampleReflectValue.Kind() == reflect.Ptr {
		sampleReflectValue = sampleReflectValue.Elem()
	}

	var sampleReflectType = sampleReflectValue.Type()

	if sampleReflectValue.NumField() > constant.TwentySix {
		err := errors.New("Headers not supported above 26")
		if err != nil {
			return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
		}
	}

	f := excelize.NewFile()
	sheet, _ := f.NewSheet(sheetName)

	// write headers column
	for i := 0; i < sampleReflectValue.NumField(); i++ {
		err := f.SetCellValue(sheetName, character[i]+"1", sampleReflectType.Field(i).Tag.Get("column"))
		if err != nil {
			logger.Error(ctx.Request.Context(), err)
			return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
		}
	}

	for i := 0; i < len(ret); i++ {
		valuesData := reflect.ValueOf(ret[i])
		if valuesData.Kind() == reflect.Ptr {
			valuesData = valuesData.Elem()
		}
		for j := 0; j < valuesData.NumField(); j++ {
			axis := fmt.Sprintf("%s%d", character[j], i+constant.Two)
			err := f.SetCellValue(sheetName, axis, valuesData.Field(j).Interface())
			if err != nil {
				return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
			}
		}
	}

	f.SetActiveSheet(sheet)

	err := f.Write(ctx.Writer)
	if err != nil {
		return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
	}

	return nil
}

// WriteExcelFromMap is a function to write data to excel file from map
// key of data must be same with headers value, the value of headers is the column name
func (e *Excelize) WriteExcelFromMap(ctx *gin.Context, sheetName string, data []map[string]interface{}, headers []string) error {
	if len(headers) > constant.TwentySix {
		err := errors.New("Headers not supported above 26")
		if err != nil {
			return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
		}
	}

	f := excelize.NewFile()
	sheet, _ := f.NewSheet(sheetName)

	// write headers column
	i := 0
	for _, h := range headers {
		err := f.SetCellValue(sheetName, character[i]+"1", h)
		i++
		if err != nil {
			logger.Error(ctx.Request.Context(), err)
			return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
		}
	}

	// write data
	// for every rows
	rowItr := 2
	for _, mapData := range data {
		colItr := 0
		// for every column
		for _, h := range headers {
			axis := fmt.Sprintf("%s%d", character[colItr], rowItr)
			colItr++
			err := f.SetCellValue(sheetName, axis, mapData[h])
			if err != nil {
				logger.Error(ctx.Request.Context(), err)
				return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
			}
		}
		rowItr++
	}

	f.SetActiveSheet(sheet)

	err := f.Write(ctx.Writer)
	if err != nil {
		return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
	}

	return nil
}

// ReadExcel is a function to read data from excel file
func (e *Excelize) ReadExcel(ctx *gin.Context, file multipart.File) ([][]string, error) {
	sheetName := "Sheet1"
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
	}

	// copy value merged cell
	mergedCells, _ := f.GetMergeCells(sheetName)
	for _, mergedCell := range mergedCells {
		err = f.UnmergeCell(sheetName, mergedCell.GetStartAxis(), mergedCell.GetEndAxis())
		if err != nil {
			return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
		}

		startX, startY, err := excelize.CellNameToCoordinates(mergedCell.GetStartAxis())
		if err != nil {
			return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
		}

		endX, endY, err := excelize.CellNameToCoordinates(mergedCell.GetEndAxis())
		if err != nil {
			return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
		}

		val := mergedCell.GetCellValue()
		for i := startY; i <= endY; i++ {
			for j := startX; j <= endX; j++ {
				axis, err := excelize.CoordinatesToCellName(j, i)
				if err != nil {
					return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
				}

				err = f.SetCellValue(sheetName, axis, val)
				if err != nil {
					return nil, errors.Wrap(err, "[Excelize-ReadExcel] error while read excel")
				}
			}
		}
	}

	rows, _ := f.GetRows(sheetName)
	return rows, nil
}

// WriteExcelMultipleSheet is a function to write data to excel file from struct
func (e *Excelize) WriteExcelMultipleSheet(ctx *gin.Context, sheetName []string, data []interface{}, headers []string) error {
	f := excelize.NewFile()
	for iSheet := 0; iSheet < len(sheetName); iSheet++ {
		reflectValue := reflect.ValueOf(data[iSheet])
		ret := make([]interface{}, reflectValue.Len())
		for i := 0; i < reflectValue.Len(); i++ {
			ret[i] = reflectValue.Index(i).Interface()
		}

		// if len(ret) == 0 {
		// 	return errors.Wrap(errors.New("Data is empty"), "[Excelize-WriteExcel] error while write excel")
		// }

		var sampleReflectValue reflect.Value
		if sampleReflectValue.Kind() == reflect.Ptr {
			sampleReflectValue = sampleReflectValue.Elem()
			logger.ErrorFromStr(ctx.Request.Context(), "sampleReflectValue: "+sampleReflectValue.String())
		}

		f.NewSheet(sheetName[iSheet])

		// write headers column
		for i := 0; i < len(headers); i++ {
			err := f.SetCellValue(sheetName[iSheet], character[i]+"1", headers[i])
			if err != nil {
				logger.ErrorWithStr(ctx.Request.Context(), "error while write excel", err)
				return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
			}
		}

		for i := 0; i < len(ret); i++ {
			valuesData := reflect.ValueOf(ret[i])
			if valuesData.Kind() == reflect.Ptr {
				valuesData = valuesData.Elem()
			}
			for j := 0; j < valuesData.NumField(); j++ {
				axis := fmt.Sprintf("%s%d", character[j], i+constant.Two)
				err := f.SetCellValue(sheetName[iSheet], axis, valuesData.Field(j).Interface())
				if err != nil {
					return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
				}
			}
		}
	}

	f.SetActiveSheet(1)
	f.DeleteSheet("Sheet1")

	if err := f.Write(ctx.Writer); err != nil {
		return errors.Wrap(err, "[Excelize-WriteExcel] error while write excel")
	}

	return nil
}
