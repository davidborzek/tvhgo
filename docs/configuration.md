tvhgo can be configured using environment variables or a config file.

> Note: Environment Variables will be overridden by the config file.

## Config via File

The configuration file in in the YAML format and looks like this:

```yaml
server:
  host: 127.0.0.1
  port: 8080

tvheadend:
  scheme: http
  host: <tvheadend_host>
  port: 9981
  username: <tvheadend_username>
  password: <tvheadend_password>

database:
  path: ./tvhgo.db
```

tvhgo will search for a config files at the following paths with the following order:

- `./config.yml`
- `./config.yaml`
- `/etc/tvhgo/config.yml`
- `/etc/tvhgo/config.yaml`

Alternatively, you can specify the path to the config file using the `TVHGO_CONFIG` environment variable
or the `--config` flag.

```bash
# Using the flag
tvhgo --config /path/to/config.yml

# Using the environment variable
TVHGO_CONFIG=/path/to/config.yml tvhgo
```

## Config via Environment Variable

To configure tvhgo using environment variables, you can use the following syntax:

```bash
TVHGO_<SECTION_NAMES...>_<CONFIG_PROPERTY>=<VALUE>
```

For example, to configure the tvheadend password:

```bash
TVHGO_TVHEADEND_PASSWORD=<YOUR_PASSWORD>
```

## Options

### Server config (server)

| Parameter  | Type       | Required | Default | Description                                                               |
| ---------- | ---------- | -------- | ------- | ------------------------------------------------------------------------- |
| host       | string     | false    |         | Bind host of the http server.                                             |
| port       | int        | false    | 8080    | Bind port of the http server. Ports below `1024` may require root rights. |
| swagger_ui | Swagger UI | false    |         | Swagger UI config.                                                        |

**Example**

```yaml
server:
  host: 127.0.0.1
  port: 1234
```

#### Swagger UI config (swagger_ui)

| Parameter | Type    | Required | Default | Description                      |
| --------- | ------- | -------- | ------- | -------------------------------- |
| enabled   | boolean | false    | true    | Enables the Swagger UI api docs. |

If enabled, you can access Swagger UI by visiting `http(s)://your-server(:8080)/api/swagger`.

**Example**

```yaml
server:
  swagger_ui:
    enabled: false
```

### Log config (log)

| Parameter | Type                                      | Required | Default | Description     |
| --------- | ----------------------------------------- | -------- | ------- | --------------- |
| level     | enum (debug, info, warning, error, fatal) | false    | info    | The log level.  |
| format    | enum(console, json)                       | false    | console | The log format. |

**Example**

```yaml
log:
  level: debug
  format: json
```

### Tvheadend config (tvheadend)

| Parameter | Type   | Required | Default | Description                                                                                      |
| --------- | ------ | -------- | ------- | ------------------------------------------------------------------------------------------------ |
| scheme    | string | false    | http    | Scheme of the tvheadend server url.                                                              |
| host      | string | **true** |         | Host (hostname or ip) of the tvheadend server.                                                   |
| port      | int    | false    | 9981    | Port of the tvheadend server.                                                                    |
| username  | string | false    |         | Username of the tvheadend server.                                                                |
| password  | string | false    |         | Password of the tvheadend server. It is recommended to configure it via an environment variable. |

> NOTE: Currently only plain text authentication is supported. Set the authentication type in tvheadent either to `Plain (insecure)` or `Both plain and digest`.

**Example**

```yaml
tvheadend:
  scheme: https
  host: 10.0.0.2
  port: 5555
  username: my_user
  password: supersecret
```

### Database config (database)

| Parameter | Type                                                                                                        | Required | Default                                           | Description                                       |
| --------- | ----------------------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------- | ------------------------------------------------- |
| type      | enum(`sqlite3`, `postgres`)                                                                                 | false    | sqlite3                                           | Type of the database.                             |
| path      | string                                                                                                      | false    | ./tvhgo.db (for the docker image: /data/tvhgo.db) | Path of the database file. (only for `sqlite3`)   |
| host      | string                                                                                                      | false    | localhost                                         | Host of the database. (only for `postgres`)       |
| port      | int                                                                                                         | false    | 5432                                              | Port of the database. (only for `postgres`)       |
| user      | string                                                                                                      | false    | tvhgo                                             | The database user. (only for `postgres`)          |
| database  | string                                                                                                      | false    | tvhgo                                             | The database name. (only for `postgres`)          |
| ssl_mode  | see [PostgresSQL Docs](https://www.postgresql.org/docs/current/libpq-ssl.html#LIBPQ-SSL-SSLMODE-STATEMENTS) | false    | disable                                           | SSL mode of the connection. (only for `postgres`) |

**Example: sqlite3**

```yaml
database:
  type: sqlite3
  path: /path/to/database.db
```

**Example: postgres**

```yaml
database:
  type: postgres
  host: 127.0.0.1
  password: supersecret
```

### Auth config (auth)

#### Session config (auth.session)

| Parameter                 | Type          | Required | Default       | Description                                                                                                           |
| ------------------------- | ------------- | -------- | ------------- | --------------------------------------------------------------------------------------------------------------------- |
| cookie_name               | string        | false    | tvhgo_session | The name of the session cookie.                                                                                       |
| cookie_secure             | bool          | false    | false         | Sets the `secure` attribute of the session cookie.                                                                    |
| maximum_inactive_lifetime | time.Duration | false    | 168h          | The maximum inactive lifetime of a session. This configures the maximum lifetime of the session after the last login. |
| maximum_lifetime          | time.Duration | false    | 720h          | The maximum lifetime of a session.                                                                                    |
| token_rotation_interval   | time.Duration | false    | 30m           | The rotation interval of the token.                                                                                   |
| cleanup_interval          | time.Duration | false    | 12h           | The interval of the scheduler to cleanup expired sessions.                                                            |

**Example**

```yaml
auth:
  session:
    cookie_name: foobar_session
    cookie_secure: true
    maximum_inactive_lifetime: 24h
    maximum_lifetime: 168h
    token_rotation_interval: 1h
    cleanup_interval: 6h
```

#### TOTP config (auth.totp)

| Parameter | Type   | Required | Default | Description      |
| --------- | ------ | -------- | ------- | ---------------- |
| issuer    | string | false    | tvhgo   | The totp issuer. |

**Example**

```yaml
auth:
  totp:
    issuer: foobar
```

#### Reverse proxy auth config (auth.reverse_proxy)

| Parameter          | Type     | Required | Default      | Description                                                              |
| ------------------ | -------- | -------- | ------------ | ------------------------------------------------------------------------ |
| enabled            | bool     | false    | false        | Enable reverse proxy authentication.                                     |
| user_header        | string   | false    | Remote-User  | The header containing the username.                                      |
| email_header       | string   | false    | Remote-Email | The header containing the email.                                         |
| name_header        | string   | false    | Remote-Name  | The header containing the name.                                          |
| allowed_proxies    | []string | false    | []           | List of allowed proxies. If not set, all requests will be blocked.       |
| allow_registration | bool     | false    | false        | If this is enabled, not existing users will automatically be registered. |

### Metrics config (metrics)

| Parameter | Type   | Required | Default  | Description                                                                  |
| --------- | ------ | -------- | -------- | ---------------------------------------------------------------------------- |
| enabled   | bool   | false    | false    | This enables the metrics endpoint exposing metrics in the prometheus format. |
| path      | string | false    | /metrics | The http path where the metrics are exposed.                                 |
| host      | string | false    |          | Bind host of the metrics http server.                                        |
| port      | int    | false    | 8081     | The http port of the metrics server.                                         |
| token     | string | false    |          | Bearer authentication token of the http metrics endpoint.                    |

> NOTE: The port cannot be the same port as the normal tvhgo http server.

**Example**

```yaml
metrics:
  enabled: true
  path: /prometheus
  host: 127.0.0.1
  port: 2345
  token: supersecret
```
