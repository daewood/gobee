package mwhttp

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/daewood/gobee/transport/httpruntime"
	"github.com/daewood/gobee/transport/middlewares/mwcommon"

	"github.com/pkg/errors"
)

// Recover recovers HTTP server from handlers' panics.
func Recover(logger interface{}) Middleware {
	logFunc := mwcommon.GetLogFunc(logger)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					stack := string(debug.Stack())
					httpruntime.SetError(
						r.Context(),
						r, w,
						errors.Errorf("recovered from panic: %v", rec),
						map[string]string{"stack": stack},
					)
					logFunc(r.Context(), fmt.Sprintf("recovered from panic: %v, %v ", rec, stack))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
