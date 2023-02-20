package auth

import (
	"net/http"
	"os"
	"strconv"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("The Authorization header is empty"))
			return
		}
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
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	})
}
