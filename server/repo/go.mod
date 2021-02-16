module repo

go 1.15

replace model => ../model

require (
	github.com/go-sql-driver/mysql v1.5.0
	model v0.0.0-00010101000000-000000000000
)
