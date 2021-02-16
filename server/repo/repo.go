package repo

import (
	"database/sql"
	"model"
	"time"
)

//	add cashflow
func AddCashflow(cashflow model.CashFlow) (bool, error) {

	var db *sql.DB
	db, err := GetMySQLDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	repository, err := GetRepo(db)
	if err != nil {
		return false, err
	}

	res, err := repository.CreateCashFlow(cashflow)

	if err != nil {
		return false, err
	}

	return res, nil

}

//	update cashflow
func UpdateCashflow(cashflow model.CashFlow) (bool, error) {

	var db *sql.DB
	db, err := GetMySQLDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	repository, err := GetRepo(db)
	if err != nil {
		return false, err
	}

	res, err := repository.UpdateCashFlow(cashflow)

	if err != nil {
		return false, err
	}

	return res, nil

}

func GetCashflow(from, to time.Time) ([]model.CashFlow, error) {
	var db *sql.DB
	db, err := GetMySQLDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository, err := GetRepo(db)
	if err != nil {
		return nil, err
	}

	res, err := repository.GetCashFlowInRange(from, to)

	if err != nil {
		return nil, err
	}

	return res, nil

}
