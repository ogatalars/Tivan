package Structs

type DocumentCNJ struct {
	Id_item   		string
	Id_item_upper   string
	Name        	string
}

type DocumentJus struct {
	Id   		string
	Doc_name   	string
	URL        	string
}

type CNJArray struct {
	Doc_name   		string
	Id_item   		string
	Id_item_upper   string
	Name        	string

}

type FinalData struct {
	Id   			string
	Doc_name   		string
	CNJFirst		string
	CnjReturn 		[]CNJArray
	URL             string
}

type NonClass struct {
	Id   			string
	Doc_name   		string
	URL             string
}
