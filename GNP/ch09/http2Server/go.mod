module http2Server

go 1.23.2

replace ch09/handlers => ../handlers

replace ch09/middleware => ../middleware

require (
	ch09/handlers v0.0.0-00010101000000-000000000000
	ch09/middleware v0.0.0-00010101000000-000000000000
)
