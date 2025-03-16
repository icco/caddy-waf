package waf

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(WAF{})
}

// WAF is a Caddy module that implements an HTTP middleware that blocks stuff.
type WAF struct {
}

// CaddyModule returns the Caddy module information.
func (WAF) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "icco.waf",
		New: func() caddy.Module { return new(WAF) },
	}
}
