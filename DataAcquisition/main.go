package main

import (
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/Acquisition"
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/DBOperation"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DBOperation.CreateDBIfNotExist()
	Acquisition.AcqAllSymbol()
	//Acquisition.AcqBySymbol("gnxusdt")
}
