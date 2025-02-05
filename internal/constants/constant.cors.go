package constants

const (
	AllowOrigin     = "*" // more specific "localhost:3000, google.com"
	AllowCredential = "true"
	AllowHeader     = "Content-Type, Connection, Sec-Fetch-Mode, Sec-Fetch-Site, Sec-Fetch-Dest, Accept-Language, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Referer, accept, Origin, Cache-Control, X-Requested-With, User-Agent, Accept, Postman-Token, Access-Control-Request-Method, Access-Control-Request-Headers" // separate with ", "
	AllowMethods    = "POST, GET, PUT, DELETE, PATCH, OPTIONS"
	MaxAge          = "43200" // for 12 hour
)
