package domain

type JwtClaims struct {
	Issuer  interface{}
	Subject interface{}
	Iat     interface{}
	Exp     interface{}
}
