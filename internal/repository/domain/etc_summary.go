package domain

type ETCSummary struct {
	Amount    int     `json:"amount"`
	Count     int     `json:"count"`
	Strip     []Strip `json:"strip"`
	PieEntry  []Pie   `json:"pieEntry"`
	PieExit   []Pie   `json:"pieExit"`
	QueryType string  `json:"queryType"`
}

type Strip struct {
	Month string `json:"entry_month"`
	Year  string `json:"entry_year"`
	Day   string `json:"entry_day"`
	Sum   int    `json:"sum"`
}

type Pie struct {
	Entry string `json:"entry"`
	Exit  string `json:"exit"`
	Sum   int    `json:"sum"`
}
