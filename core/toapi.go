package core

import (
	"fmt"
	"log"

	"github.com/ledongthuc/goterators"
	"github.com/xuri/excelize/v2"
)

func GetRows(filePath string, sheetName string) ([]map[string]interface{}, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("fail to open file '%s': %w", filePath, err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Printf("fail to close file '%s': %v\n", filePath, err)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("fail to get rows from file '%s' of sheet '%s': %w", filePath, sheetName, err)
	}

	if len(rows) == 0 {
		return []map[string]interface{}{}, nil
	}

	header := getHeader(rows[0])

	returningRows := make([]map[string]interface{}, 0, len(rows))
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}
		isEmptyRow := goterators.Every(row, func(cellValue string) bool {
			return len(cellValue) == 0
		})
		if isEmptyRow {
			continue
		}

		returningRow := map[string]interface{}{}
		for columnIndex, headerName := range header {
			returningRow[headerName] = row[columnIndex]
		}
		returningRows = append(returningRows, returningRow)
	}
	return returningRows, nil
}

func getHeader(firstCells []string) map[int64]string {
	header := map[int64]string{}
	for index, cell := range firstCells {
		if len(cell) > 0 {
			header[int64(index)] = cell
		}
	}
	return header
}
