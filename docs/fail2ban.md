Fail2Ban will prevent attackers to brute force your tvhgo instance. This is important if your server is publicly available.

## Filter

Create a new filter file at `/etc/fail2ban/filter.d/tvhgo.local`:

```ini
[Definition]
failregex =  .*login failed: invalid username or password.*ip=<HOST> username=.*
ignoreregex =
```

## Jail

Create a new jail file at `/etc/fail2ban/jail.d/tvhgo.local`:

```ini
[tvhgo]
enabled = true
filter = tvhgo
maxretry = 3
bantime = 14400
findtime = 14400
action = iptables-allports[chain="INPUT"]
```

**Info for docker users**

For docker you have to use the FORWARD chain instead of the INPUT chain:

```ini
action = iptables-allports[chain="FORWARD"]
```
