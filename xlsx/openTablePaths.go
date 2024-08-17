package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

func OpenTablePaths(path1, path2 string, log *zap.Logger) (*excelize.File, *excelize.File, error) {
	if path1 == "" || path2 == "" {
		return nil, nil, fmt.Errorf("cannot open tables, path is empty")
	}

	f1, err := excelize.OpenFile(path1)
	if err != nil {
		log.Error(err.Error())
		return nil, nil, fmt.Errorf("cannot open %s with error: %v", path1, err)
	}

	f2, err := excelize.OpenFile(path2)
	if err != nil {
		log.Error(err.Error())
		return nil, nil, fmt.Errorf("cannot open %s with error: %v", path2, err)
	}

	return f1, f2, nil
}
