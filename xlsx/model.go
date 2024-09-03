package xlsx

import (
	"expense_tables/config"

	"go.uber.org/zap"
)

type ProductsInfo struct {
	Log         *zap.Logger
	Config      *config.Config
	Products1   map[string]int
	Products2   map[string]int
	ProductsSum map[string]int
	SumPrice    int
}

func NewInfo(log *zap.Logger, cfg *config.Config) *ProductsInfo {
	return &ProductsInfo{
		Log:         log,
		Config:      cfg,
		Products1:   make(map[string]int),
		Products2:   make(map[string]int),
		ProductsSum: make(map[string]int),
	}
}
