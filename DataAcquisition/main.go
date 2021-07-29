package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kuianchen/Lasagna/DataAcquisition/DBOperation"
)

func main() {
	DBOperation.CreateDBIfNotExist()

}
