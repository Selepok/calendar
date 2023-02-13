package auth

import (
	"net/http"
	"os"
	"strconv"
)

//type Middleware struct {
//}
// auth/auth.go

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expirationMinutes, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRATION_TIME_IN_MINUTES"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jwt := JwtWrapper{
			SecretKey:         os.Getenv("SECRET_KEY"),
			ExpirationMinutes: expirationMinutes,
		}

		err = jwt.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
