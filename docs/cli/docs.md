# NAME

tvhgo - Modern and secure api and web interface for Tvheadend

# SYNOPSIS

tvhgo

```
[--config|-c]=[value]
```

**Usage**:

```
tvhgo [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--config, -c**="": Path to the configuration file


# COMMANDS

## server

Starts the tvhgo server

## admin

Admin controls for the tvhgo server

### user

Manage users of the tvhgo server

#### add

Add a new user

**--display-name, -n**="": Display name of the new user

**--email, -e**="": Email of the new user

**--password, -p**="": Password of the new user

**--username, -u**="": Username of the new user

#### list

List users.

#### delete

Deletes a user

**--username, -u**="": Username of the new user

#### 2fa

Manage 2FA of a user

##### disable

Disable 2FA for a user.

**--username, -u**="": Username of the new user

#### token

Manage tokens of a user

##### list

List tokens of a user

**--username, -u**="": Username of the new user

##### generate

Generate a new token

**--name, -n**="": Name of the token,

**--username, -u**="": Username of the new user

##### revoke

Revokes a token

**--id**="": ID of the token,
