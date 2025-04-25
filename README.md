Autodns module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with autodns.

## Caddy module name

```
dns.providers.autodns
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "autodns",
				"username": "{env.AUTODNS_USERNAME}",
				"password": "{env.AUTODNS_PASSWORD}",
				"endpoint": "{env.AUTODNS_ENDPOINT}",
				"context": "{env.AUTODNS_CONTEXT}"
			}
		}
	}
}
```

### Caddyfile Configuration

Caddyfile configuration

#### Global Configuration

```
{
  acme_dns autodns {
    username {env.AUTODNS_USERNAME}
		password {env.AUTODNS_PASSWORD}
		endpoint {env.AUTODNS_ENDPOINT}
		context {env.AUTODNS_CONTEXT}
  }
}
```

#### Per-Site Configuration

```
tls {
	dns autodns {
		username {env.AUTODNS_USERNAME}
		password {env.AUTODNS_PASSWORD}
		endpoint {env.AUTODNS_ENDPOINT}
		context {env.AUTODNS_CONTEXT}
	}
}
```


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/autodns) for important information about credentials.
