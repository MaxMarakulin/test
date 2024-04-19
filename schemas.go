package main

type exchangeRateModel struct {
	Id     int32   `gorm:"primaryKey"`
	To     int32   `gorm:"column:to"`
	From   int32   `gorm:"column:from"`
	Value  float64 `gorm:"column:value"`
	Source string  `gorm:"column:source"`
}

type response struct {
	Id     int32   `gorm:"column:id"`
	To     string  `gorm:"column:to"`
	From   string  `gorm:"column:from"`
	Value  float64 `gorm:"column:value"`
	Source string  `gorm:"column:source"`
}

type currenciesModel struct {
	Id   int32  `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

type rateCBR struct {
	Currency currency `json:"Valute"`
}

type currency struct {
	BYN currencyModel `json:"BYN"`
}

type currencyModel struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

type rateNBRB struct {
	CurID           int     `json:"Cur_ID"`
	Date            string  `json:"Date"`
	CurAbbreviation string  `json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `json:"Cur_Name"`
	CurOfficialRate float64 `json:"Cur_OfficialRate"`
}
