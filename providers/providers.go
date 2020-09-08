package providers

// TokenAuthProvider allows to get request signed using generated JWT
type TokenAuthProvider interface {
	GetToken(audience string) (string, error)
}
