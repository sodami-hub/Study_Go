module main

go 1.23.2

replace housework => ../housework

replace json => ../json

replace gob => ../gob

require (
	gob v0.0.0-00010101000000-000000000000
	housework v0.0.0-00010101000000-000000000000
)
