package models

type Part struct {
	Id		int		`json:"id,omitempty"`
	Name		string		`json:"name,omitempty"`
	Desc		string		`json:"desc,omitempty"`
	Qty		int		`json:"qty,omitempty"`
}

type Project struct {
	Id		int		`json:"id,omitempty"`
	Name		string		`json:"name"`
	Desc		string		`json:"desc"`
	ImagePath	string		`json:"imgPath"`
	Url		string		`json:"url"`
	Parts		[]*Part		`json:"parts"`
}
