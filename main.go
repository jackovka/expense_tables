package main

import (
	"expense_tables/config"
	"expense_tables/logger"
	xlsx "expense_tables/xlsx"
	"fmt"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("ERROR: cannot read config with error: %v\n", err)
		return
	}

	log := logger.NewLogger(cfg)
	log.Info("Start application")

	f1, f2, err := xlsx.OpenTablePaths(cfg.TablePath1, cfg.TablePath2, log)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Debug("Success opened tables")
	defer f1.Close()
	defer f2.Close()

	info := xlsx.NewInfo(log, cfg)

	err = info.GetProducts(f1, f2)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// fmt.Println(info.Products1)
	// fmt.Println(info.Products2)
	// fmt.Println(info.ProductsSum)

	if cfg.CountUsers == 1 {
		info.JoinTablesUser()
	} else if cfg.CountUsers == 2 {
		info.JoinTablesUsers()
	} else {
		log.Error("too much users!")
	}

	log.Info("Success joined tables")
}
