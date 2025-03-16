# caddy-waf

A dumb and aggressive firewall for Caddy v2.

## Sources

The following are data sources for bad actors:

 - https://github.com/crowdsecurity/crowdsec
 - https://github.com/The-Shadowserver-Foundation/api_utils
 - https://owasp.org/www-project-modsecurity-core-rule-set/ 
 - https://github.com/corazawaf/coraza

## Design Goals

 - Be as aggressive as possible
 - Be as simple as possible
 - Be as fast as possible

This WAF is designed for home users and small businesses. It is not designed for large scale deployments. The idea is to block as many bad actors as possible with as little configuration as possible.
