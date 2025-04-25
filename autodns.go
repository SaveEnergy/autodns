package autodns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/autodns"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *autodns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.autodns",
		New: func() caddy.Module { return &Provider{new(autodns.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Username = caddy.NewReplacer().ReplaceAll(p.Provider.Username, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	p.Provider.Endpoint = caddy.NewReplacer().ReplaceAll(p.Provider.Endpoint, "")
	p.Provider.Context = caddy.NewReplacer().ReplaceAll(p.Provider.Context, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	autodns {
//	    secret_key string
//	    organization_id string
//	}
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "username":
				if p.Provider.Username != "" {
					return d.Err("Username already set")
				}
				if d.NextArg() {
					p.Provider.Username = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("Organization ID already set")
				}
				if d.NextArg() {
					p.Provider.Password = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "Endpoint":
				if p.Provider.Endpoint != "" {
					return d.Err("Endpoint already set")
				}
				if d.NextArg() {
					p.Provider.Endpoint = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "context":
				if p.Provider.Context != "" {
					return d.Err("Context already set")
				}
				if d.NextArg() {
					p.Provider.Context = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Username == "" {
		return d.Err("missing Secret key")
	}
	if p.Provider.Password == "" {
		return d.Err("missing Organization ID")
	}
	if p.Provider.Endpoint == "" {
		return d.Err("missing Endpoint")
	}
	if p.Provider.Context == "" {
		return d.Err("missing Context")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
