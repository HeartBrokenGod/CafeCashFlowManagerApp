module server

go 1.15

replace model => ./model

replace repo => ./repo

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	model v0.0.0-00010101000000-000000000000
	repo v0.0.0-00010101000000-000000000000
)
