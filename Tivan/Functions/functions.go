package Functions

import (
	"Tivan/Structs"
	"strings"
)

func splitName(docName string) string{
	text := strings.Split(docName, "-")
	if len(text) > 1 {

		text1 := strings.Split(text[1], ".")

		if len(text1) > 1 {

			return newName(text1[0])

		}else{

			return newName(text[1])

		}

	}else{

		return newName(docName)

	}
}

func returnCNJ(docNameSplit string,data []Structs.DocumentCNJ) ([]Structs.CNJArray, string,bool){
	var dataCNJ []Structs.CNJArray
	var cnjFirst string
	flag := false

	for i := 0; i < len(data); i++ {

		if strings.Contains(docNameSplit, data[i].Name) {

			dataCNJ = append(dataCNJ, Structs.CNJArray{
				Doc_name:      docNameSplit,
				Id_item:       data[i].Id_item,
				Id_item_upper: data[i].Id_item_upper,
				Name:          data[i].Name,
			})

			flag = true
		}
	}
	if flag == true {
		cnjFirst = dataCNJ[0].Name
	}
	return  dataCNJ, cnjFirst, flag
}

func newName(docName string) string{
	var nameReturn string

	if docName == "Petições Diversas" || docName == "PETIÇÕES DIVERSAS" {

		nameReturn = "Petição (Outras)"

	} else if docName == "Petições Intermediárias Diversas" {

		nameReturn = "Petição Intermediária"

	} else if docName == " Documento Diverso"{

		nameReturn = "Documentos Diversos"

	} else if docName == " Solicitação de Habilitação"{

		nameReturn = "Pedido de Habilitação"

	} else if docName == " Apresentação de Substabelecimento com Reserva de Poderes" || docName == " Substabelecimento com Reserva de Poderes" || docName == " Apresentação de Substabelecimento sem Reserva de Poderes" || docName == " Substabelecimento sem Reserva de Poderes" || docName == " Instrumento de Mandato  Apresentação de Substabelecimento com reservas de poderes" || docName == " Instrumento de Mandato  Apresentação de Substabelecimento sem reservas de poderes" || docName == " Substabelecimento com Reserva de Poderes (inativo)" {

		nameReturn = "Instrumento de procuração"

	} else if docName == " Procuração" || docName == " Apresentação de Procuração" || docName == " Instrumento de Mandato  Apresentação de Procuração" {

		nameReturn = "Instrumento de procuração"

	} else if docName == " Endereço Localizado" || docName == " Endereço  Apresentação" {

		nameReturn = "Comprovante de Endereço"

	} else if docName == " Indicação de Bens à Penhora" || docName == " Penhora Indicação de Bens" || docName == " Penhora  Indicação de Bens" || docName == " Outros bens ou direitos" {

		nameReturn = "Nomeação de Bens à Penhora"

	} else if docName == " Requerimento de Adiamento de Audiência" || docName == " Audiência  Requerimento de Designação" || docName == " Audiência  Requerimento de Adiamento" || docName == " Audiência  Requerimento de Antecipação"  {

		nameReturn = "Pedido de Designação/Redesignação de Audiência"

	} else if docName == " Audiência" || docName == " Audiência Realizada Exitosa"{

		nameReturn = "Ata de Audiência"

	} else if docName == " Quesitos" || docName == " Quesitos Complementares"{

		nameReturn = "Apresentação de Quesitos/Indicação de Assistente Técnico"

	} else if docName == " Requisição Final de Honorários Periciais" || docName == " Apresentação de Proposta de Honorários Periciais"{

		nameReturn = "Solicitação de Honorários de Perito"

	} else if docName == " Solicitação do Perito" || docName == " Solicitação do Perito (Digitalizado)"{

		nameReturn = "Esclarecimento de Perito"

	} else if docName == " Peticionamento Eletrônico  Petição Peritos" {

		nameReturn = "Manifestação do Perito"

	} else if docName == " Prova Pericial  Art 474 do CPC" {

		nameReturn = "Agendamento de Vistoria - Prova Pericial - Art. 474 do CPC"

	} else{

		nameReturn = docName

	}


	return  nameReturn
}