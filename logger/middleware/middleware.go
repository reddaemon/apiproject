package middleware

import (
	"log"
	"net/http"

	"github.com/reddaemon/apiproject/config"
	lh "github.com/reddaemon/apiproject/handlers"
	"github.com/reddaemon/apiproject/logger"
)

//LogRequestHandler logging all requests
func LogRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := config.GetConfig(*lh.ConfigPath)
		l, err := logger.GetLogger(c)
		if err != nil {
			log.Fatalf("Cannot get logger %v", err)
		}
		lsugar := l.Sugar()
		//log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		lsugar.Infof("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
