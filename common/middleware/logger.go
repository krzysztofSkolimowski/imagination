package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	HTTPStartingMsg = "Starting HTTP request"
	HTTPDoneMsg     = "HTTP request done"
)

// todo - create abstraction for logger in common
func RequestLogger(logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			addr := r.RemoteAddr
			if i := strings.LastIndex(addr, ":"); i != -1 {
				addr = addr[:i]
			}

			wrapWriter := &wrapResponseWriter{ResponseWriter: w}
			next.ServeHTTP(wrapWriter, r)

			defer func() {
				logger.WithFields(
					log.Fields{
						"address":   addr,
						"time":      t1.Format("Mon, 02 Jan 2006 15:04:05 MST"),
						"proto":     r.Proto,
						"status":    wrapWriter.status,
						"written":   wrapWriter.written,
						"Referer":   r.Referer(),
						"UserAgent": r.UserAgent(),
					}).Info("REQUEST")
			}()
		}

		return http.HandlerFunc(fn)
	}
}

type wrapResponseWriter struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *wrapResponseWriter) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *wrapResponseWriter) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}
