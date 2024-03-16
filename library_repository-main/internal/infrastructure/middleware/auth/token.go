package auth

import (
	"github.com/go-chi/jwtauth"
	"net/http"
)

type IToken interface {
	GetAlgorithm() string
	Encode(claims map[string]interface{}) string
	Verifier() func(http.Handler) http.Handler
	Authenticator(next http.Handler) http.Handler
}

type Token struct {
	token     *jwtauth.JWTAuth
	algorithm string
	key       string
}

func NewToken(algorithm, key string) *Token {
	return &Token{
		algorithm: algorithm,
		key:       key,
		token:     jwtauth.New(algorithm, []byte(key), nil),
	}
}

func (t *Token) GetAlgorithm() string {
	return t.algorithm
}

func (t *Token) Encode(claims map[string]interface{}) string {
	_, token, _ := t.token.Encode(claims)
	return token
}

func (t *Token) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(t.token)
}

func (t *Token) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if username, ok := claims["username"].(string); ok && username != "" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})

}
