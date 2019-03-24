package models

type Part struct {
	Id		int		`json:"id,omitempty"`
	Name		string		`json:"name"`
	Desc		string		`json:"desc"`
	Qty		int		`json:"qty,omitempty"`
}

type Project struct {
	Id		int		`json:"id,omitempty"`
	Name		string		`json:"name"`
	Desc		string		`json:"desc"`
	ImagePath	string		`json:"imgPath"`
	Parts		[]Part		`json:"parts"`
}
