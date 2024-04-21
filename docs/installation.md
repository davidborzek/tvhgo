This guide provides step-by-step instructions to install thvgo using docker.

## Prerequisites

Before you start, make sure you have the following ready:

- Docker

## Install tvhgo

You can use the following command to run tvhgo using docker.

```bash
$ docker run -d \
    --name tvhgo \
    -p 8080:8080 \
    -e 'TVHGO_TVHEADEND_HOST=<TVHEADEND_HOST>' \
    -v /path/to/store/data:/tvhgo \
    --restart unless-stopped \
    ghcr.io/davidborzek/tvhgo:latest
```

> Note: tvhgo runs as a non-root user inside the container. Make sure the mounted directory is writable by the user with UID 1000 (default).

Replace `<TVHEADEND_HOST>` with the actual hostname or ip of your tvheadend server and adapt /path/to/store/data to a path of your choice to persist the data stored by thvgo.

You can find more configuration options [here](#configuration).

## Create a user

To complete the setup you need to create a user.

```bash
docker exec -it tvhgo tvhgo admin user add
```

Follow the interactive setup to create a new user.
