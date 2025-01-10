package constants

const (
	AllowOrigin     = "*" // more specific "localhost:3000, google.com"
	AllowCredential = "true"
	AllowHeader     = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Referer, accept, origin, Cache-Control, X-Requested-With, User-Agent, Accept, Postman-Token" // separate with ", "
	AllowMethods    = "POST, GET, PUT, DELETE, PATCH"
	MaxAge          = "43200" // for 12 hour
)
