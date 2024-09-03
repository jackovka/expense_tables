package xlsx

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	sheet = "Лист1"
)

func (info *ProductsInfo) JoinTablesUsers() error {
	info.Log.Debug("Start JoinTablesUsers")

	f := excelize.NewFile()
	index, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	defer f.Close()

	// упорядочивание в алфавитном порядке
	products := make([]string, 0, len(info.ProductsSum))
	for k := range info.ProductsSum {
		products = append(products, k)
	}
	sort.Strings(products)

	f.SetColWidth(sheet, "A", "A", float64(25))
	f.SetColWidth(sheet, "B", "C", float64(10))
	f.SetColWidth(sheet, "E", "E", float64(15))

	// 1ый юзер и заполнение поля с продуктами
	f.SetCellValue(sheet, "A1", "Продукт")
	f.SetCellValue(sheet, "B1", "Траты 1")
	for i, productName := range products {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), strings.Title(productName))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), info.Products1[productName])
	}

	// 2ой юзер
	f.SetCellValue(sheet, "C1", "Траты 2")
	for i, productName := range products {
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), info.Products2[productName])
	}
	// суммарная трата
	f.SetCellValue(sheet, "E1", "Сумма трат")
	for i, productName := range products {
		f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), info.ProductsSum[productName])
	}

	f.SetCellValue(sheet, "G2", "Всего:")
	f.SetCellValue(sheet, "G3", info.SumPrice)

	// Сохранение файла
	f.SetActiveSheet(index)
	err = f.SaveAs("./tablesDoc/итог.xlsx")
	if err != nil {
		println(err.Error())
	}

	return nil
}

func (info *ProductsInfo) JoinTablesUser() error {
	info.Log.Debug("Start JoinTablesUser")

	f := excelize.NewFile()
	index, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	defer f.Close()

	// упорядочивание в алфавитном порядке
	products := make([]string, 0, len(info.ProductsSum))
	for k := range info.ProductsSum {
		products = append(products, k)
	}
	sort.Strings(products)

	f.SetColWidth(sheet, "A", "A", float64(25))
	f.SetColWidth(sheet, "B", "B", float64(13))

	f.SetCellValue(sheet, "A1", "Продукт")
	f.SetCellValue(sheet, "B1", "Сумма трат")

	f.SetCellValue(sheet, "D2", "Всего:")
	f.SetCellValue(sheet, "D3", info.SumPrice)

	for i, productName := range products {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), strings.Title(productName))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), info.ProductsSum[productName])
	}

	// Сохранение файла
	f.SetActiveSheet(index)
	err = f.SaveAs("./tablesDoc/траты.xlsx")
	if err != nil {
		println(err.Error())
	}

	return nil
}
