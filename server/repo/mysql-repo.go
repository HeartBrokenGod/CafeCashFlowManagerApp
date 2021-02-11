package repo

import (
	"database/sql"
	"model"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func GetRepo(db *sql.DB) (repo *Repo, err error) {
	return &Repo{db}, nil
}

func (repo Repo) GetCashFlowInRange(fromDate time.Time, toDate time.Time) ([]model.CashFlow, error) {
	rows, err := repo.Db.Query(`SELECT date,cash_inflow,e_pay_inflow,cash_expenditure,created_on,updated_on FROM cash_flow WHERE date BETWEEN  ?  AND  ?  `, fromDate, toDate)

	if err != nil {
		return nil, err
	}
	cashFlows := []model.CashFlow{}

	for rows.Next() {
		var date string
		var cashinflow int
		var epayinflow int
		var cashexpenditure int
		var createdon string
		var updatedon string
		err2 := rows.Scan(&date, &cashinflow, &epayinflow, &cashexpenditure, &createdon, &updatedon)

		if err2 != nil {
			return nil, err2
		}
		cashFlow := model.CashFlow{date, cashinflow, epayinflow, cashexpenditure, createdon, updatedon}
		cashFlows = append(cashFlows, cashFlow)

	}

	return cashFlows, nil

}

func (repo Repo) CreateCashFlow(cashflow model.CashFlow) (bool, error) {
	_, err := repo.Db.Exec(`INSERT INTO cash_flow (date, cash_inflow, e_pay_inflow, cash_expenditure) VALUES (?, ?, ?, ?)`, cashflow.Date, cashflow.CashInflow, cashflow.EPayInflow, cashflow.CashExpenditure)
	if err != nil {
		return false, err
	}
	return true, nil
}
