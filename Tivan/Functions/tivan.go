package Functions

import (
	"Tivan/CSV"
	"Tivan/Structs"
	"fmt"
)

const pathCNJ = "CNJ_doc/DataCNJ.csv"

func Tivan(dataJus []Structs.DocumentJus) ([]Structs.FinalData, []Structs.NonClass){
	var finalData  []Structs.FinalData
	var finalError []Structs.NonClass

	dataCNJ := CSV.ReadCSVcnj(pathCNJ)

	fmt.Println("Taneleer Tivan, the Collector: It's quite a collection!", len(dataJus))
	fmt.Println("Let's catalog ...")

	for i := 1; i < len(dataJus); i++ {

		docName := splitName(dataJus[i].Doc_name)

		cnjReturn, cnjFirst,flag := returnCNJ(docName,dataCNJ)

		if flag {

			finalData = append(finalData, Structs.FinalData{
				Id:        dataJus[i].Id,
				Doc_name:  dataJus[i].Doc_name,
				CNJFirst:  cnjFirst,
				CnjReturn: cnjReturn,
				URL: 	   dataJus[i].URL,
			})

		}else{

			finalError = append(finalError, Structs.NonClass{
				Id:       dataJus[i].Id,
				Doc_name: dataJus[i].Doc_name,
				URL: 	   dataJus[i].URL,
			})

		}

	}

	fmt.Println("Beautiful, beyond compare! What a collection ... ")

	return finalData, finalError
}

