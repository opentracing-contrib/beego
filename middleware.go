package apm

import (
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"net/url"
)

func Middleware(componentName string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {

		opts := []nethttp.MWOption{
			nethttp.MWComponentName(componentName),
			nethttp.OperationNameFunc(func(r *http.Request) string {
				return "HTTP " + r.Method + " " + r.URL.Path
			}),
			nethttp.MWURLTagFunc(func(u *url.URL) string {
				return u.String()
			}),
			nethttp.MWSpanFilter(func(r *http.Request) bool {
				return true
			}),
			nethttp.MWSpanObserver(func(span opentracing.Span, r *http.Request) {

			}),
		}

		return nethttp.Middleware(opentracing.GlobalTracer(), h, opts...)
	}
}
