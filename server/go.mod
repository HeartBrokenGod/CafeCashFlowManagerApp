module server

go 1.15

replace model => ./model

replace repo => ./repo

require (
	model v0.0.0-00010101000000-000000000000
	repo v0.0.0-00010101000000-000000000000
)
