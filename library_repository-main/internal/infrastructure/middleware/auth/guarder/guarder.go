package guarder

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/middleware/auth"
	"net/http"
	"os"
)

type Guarder interface {
	Verifier() func(http.Handler) http.Handler
	Authenticator(next http.Handler) http.Handler
	SetToken(token auth.IToken)
	Encode(claims map[string]interface{}) string
}

type Guard struct {
	token auth.IToken
}

type Option func(*Guard)

func NewGuarder(opts ...Option) Guarder {
	return newGuard(opts...)
}

func newGuard(opts ...Option) *Guard {
	guard := &Guard{
		token: auth.NewToken(os.Getenv("AUTH_ALGORITHM"), os.Getenv("AUTH_SECRET_KEY")),
	}
	for _, opt := range opts {
		opt(guard)
	}
	return guard
}

func WithToken(token auth.IToken) Option {
	return func(guard *Guard) {
		guard.token = token
	}
}

func (g *Guard) Verifier() func(http.Handler) http.Handler {
	return g.token.Verifier()
}

func (g *Guard) Authenticator(next http.Handler) http.Handler {
	return g.token.Authenticator(next)
}

func (g *Guard) SetToken(token auth.IToken) {
	g.token = token
}

func (g *Guard) Encode(claims map[string]interface{}) string {
	return g.token.Encode(claims)
}
