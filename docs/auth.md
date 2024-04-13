## Reverse proxy auth

You configure tvhgo to let a HTTP reverse proxy handle the authentication.

```yaml
auth:
  reverse_proxy:
    # Defaults to false, set to true to enable reverse proxy authentication.
    enabled: true
    # HTTP header containing the username.
    user_header: X-MY-USER_HEADER
    # HTTP header containing the email. (not required)
    # If not available the username will be used.
    email_header: X-MY-EMAIL_HEADER
    # HTTP header containing the name. (not required)
    # If not available the username will be used.
    name_header: X-MY-NAME_HEADER
    # Limit where the reverse proxy is allowed to come from to prevent
    # spoofing the headers. If not set, all requests will be blocked.
    allowed_proxies: ["192.168.1.1", "192.168.2.0/24"]
    # If this is enabled, not existing users will automatically be registered.
    allow_registration: true
```

See [Reverse proxy auth](configuration.md/#reverse-proxy-auth-config-authreverse_proxy) for further information.
