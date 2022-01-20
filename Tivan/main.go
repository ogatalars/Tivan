package main

import (
	"Tivan/CSV"
	"Tivan/Functions"
)

const pathJUS = "CNJ_doc/Query.csv"

func main() {

	dataJus := CSV.ReadCSVjus(pathJUS)

	fileOK, fileErrors := Functions.Tivan(dataJus)

	CSV.ExportCSV(fileOK,fileErrors)

}

