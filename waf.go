package waf

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyhttp"
)

// Interface guards
var (
	_ caddy.Provisioner           = (*WAF)(nil)
	_ caddy.Validator             = (*WAF)(nil)
	_ caddyhttp.MiddlewareHandler = (*WAF)(nil)
	_ caddyfile.Unmarshaler       = (*WAF)(nil)
)

func init() {
	caddy.RegisterModule(WAF{})
}

// WAF is a Caddy module that implements an HTTP middleware that blocks stuff.
type WAF struct {
}

// CaddyModule returns the Caddy module information.
func (WAF) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.icco.waf",
		New: func() caddy.Module { return new(WAF) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m WAF) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	return next.ServeHTTP(w, r)
}
