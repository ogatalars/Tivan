package CSV

import (
	"Tivan/Error"
	"Tivan/Structs"
	"encoding/csv"
	"os"
)

func ReadCSVcnj(filePath string) []Structs.DocumentCNJ{
	var data []Structs.DocumentCNJ

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer csvFile.Close()

	csvLines, err1 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err1)

	for _, line := range csvLines {
		emp := Structs.DocumentCNJ{
			Id_item: line[0],
			Id_item_upper: line[1],
			Name: line[3],
		}
		data = append(data, emp)
	}

	return data
}

func ReadCSVjus(filePath string) []Structs.DocumentJus{
	var data []Structs.DocumentJus

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer csvFile.Close()

	csvLines, err1 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err1)

	for _, line := range csvLines {
		emp := Structs.DocumentJus{
			Id:       line[0],
			Doc_name: line[1],
			URL:      line[2],
		}
		data = append(data, emp)
	}

	return data
}
