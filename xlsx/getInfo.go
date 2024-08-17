package xlsx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func (info *ProductsInfo) GetProducts(f1, f2 *excelize.File) error {
	info.Log.Debug("Start GetProducts")

	rows1, err := f1.GetRows("Лист1")
	if err != nil {
		info.Log.Error(err.Error())
		return err
	}
	info.fillProducts(rows1, 1)

	rows2, err := f2.GetRows("Лист1")
	if err != nil {
		info.Log.Error(err.Error())
		return err
	}
	info.fillProducts(rows2, 2)

	return nil
}

func (info *ProductsInfo) fillProducts(rows [][]string, userIdx int) {
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 2 {
			continue
		}

		cost, err := strconv.Atoi(row[1])
		if err != nil {
			info.Log.Error(fmt.Sprintf("cannot convert price %v: %v", row[1], err))
			continue
		}

		info.ProductsSum[strings.ToLower(row[0])] += cost
		if userIdx == 1 {
			info.Products1[strings.ToLower(row[0])] += cost
		} else {
			info.Products2[strings.ToLower(row[0])] += cost
		}
	}

	info.Log.Debug(fmt.Sprintf("Success fill products to user %d", userIdx))
}
