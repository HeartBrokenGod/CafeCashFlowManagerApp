package main

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"repo"
	"time"
)

func main() {

	// Handle func for  Adding New cashflow record
	http.HandleFunc("/api/cashflow/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var cashflow model.CashFlow
		json.NewDecoder(r.Body).Decode(&cashflow)

		_, err := repo.AddCashflow(cashflow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("ADDED cash_flow : %+v\n", cashflow)
		//	code to log file
		//	write Response
		json.NewEncoder(w).Encode(cashflow)
		return

	})

	// Handle func for  Updating cashflow record
	http.HandleFunc("/api/cashflow/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var cashflow model.CashFlow
		json.NewDecoder(r.Body).Decode(&cashflow)

		_, err := repo.UpdateCashflow(cashflow)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("UPDATED cash_flow : %+v\n", cashflow)
		//	code to log file
		//	write Response
		json.NewEncoder(w).Encode(cashflow)
		return

	})

	// Handle func for Reading cashflow record
	http.HandleFunc("/api/cashflow/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		type DateRange struct {
			From string `json="From"`
			To   string `json:"To"`
		}

		var dateRange DateRange

		json.NewDecoder(r.Body).Decode(&dateRange)

		fromDate, err := time.Parse("2006-01-02", dateRange.From)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		toDate, err := time.Parse("2006-01-02", dateRange.To)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cashflows, err := repo.GetCashflow(fromDate, toDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("READ cash_flow : %+v\n", dateRange)
		// code to log file
		// write Response
		json.NewEncoder(w).Encode(cashflows)
		return

	})

	http.ListenAndServe(":8080", nil)

}
