package CSV

import (
	"Tivan/Error"
	"Tivan/Structs"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	fmt.Println("CSV - ready")
	return os.Create(p)
}

func exportCSVok(data []Structs.FinalData){
	var empData [][]string

	nameCSV := "Docs_CNJ_OK"

	head := []string{"id","doc_name", "CNJfirst","CnjReturn", "URL"}
	empData = append(empData, head)

	for i := 0; i < len(data); i++ {
		cnjReturn := ""
		for  j := 0; j < len(data[i].CnjReturn); j++ {
			cnjReturn += "{"+data[i].CnjReturn[j].Doc_name+"," +data[i].CnjReturn[j].Id_item+"," +data[i].CnjReturn[j].Id_item_upper+"," +data[i].CnjReturn[j].Name+" } "
		}
		final := []string{
			data[i].Id,
			data[i].Doc_name,
			data[i].CNJFirst,
			cnjReturn,
			data[i].URL,
		}
		empData = append(empData, final)
	}

	csvFile, _ := createFile("Return Files/" + nameCSV + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()

	err := csvFile.Close()
	Error.CheckError(err)

}

func exportCSVerror(data []Structs.NonClass){
	var empData [][]string

	nameCSV := "Docs_CNJ_error"

	head := []string{"id","doc_name", "URL"}
	empData = append(empData, head)

	for i := 0; i < len(data); i++ {
		final := []string{
			data[i].Id,
			data[i].Doc_name,
			data[i].URL,
		}
		empData = append(empData, final)
	}

	csvFile, _ := createFile("Return Files/" + nameCSV + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()

	err := csvFile.Close()
	Error.CheckError(err)

}

func ExportCSV(dataOK []Structs.FinalData, dataError []Structs.NonClass){

	fmt.Println("Total CNJ found:",len(dataOK))
	fmt.Println("Total not found:", len(dataError))

	exportCSVerror(dataError)
	exportCSVok(dataOK)

}
