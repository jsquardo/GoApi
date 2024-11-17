package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.new("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

	})
}
