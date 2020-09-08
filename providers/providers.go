package providers

// JWTAuthProvider allows to get request signed using generated JWT
type JWTAuthProvider interface {
	GetJWT(audience string) (string, error)
}
