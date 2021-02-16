package model

type CashFlow struct {
	Date            string `json:"Date"`
	CashInflow      int    `json:"CashInflow"`
	EPayInflow      int    `json:"EPayInflow"`
	CashExpenditure int    `json:"CashExpenditure"`
	CreatedOn       string `json:"CreatedOn"`
	UpdatedOn       string `json:"UpdatedOn"`
}
