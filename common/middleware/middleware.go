package middleware

import "net/http"

type MiddlewareFn func(next http.Handler) http.Handler